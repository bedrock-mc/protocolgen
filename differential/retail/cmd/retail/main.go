// Command retail prepares the retail differential inputs, runs the Go
// decoders over them and compares every decoder's verdicts.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"protocolgen/differential/retail"
)

const usage = `usage: retail <command> [flags]

  prepare  verify the pinned fixtures and write every input
  decode   run the generated Go codec and gophertunnel over the inputs
  compare  compare all four decoders' verdicts against the exceptions`

var errFailed = errors.New("comparison failed")

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
	commands := map[string]func([]string) error{"prepare": prepare, "decode": decode, "compare": compare}
	command, ok := commands[os.Args[1]]
	if !ok {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
	if err := command(os.Args[2:]); err != nil {
		if !errors.Is(err, errFailed) {
			fmt.Fprintln(os.Stderr, "retail:", err)
		}
		os.Exit(1)
	}
}

func prepare(args []string) error {
	flags := flag.NewFlagSet("prepare", flag.ExitOnError)
	corpusPath := flags.String("corpus", "corpus.json", "retail corpus")
	cinnabar := flags.String("cinnabar", "", "Cinnabar checkout at the corpus commit")
	out := flags.String("out", "", "output directory")
	flags.Parse(args)
	if *cinnabar == "" || *out == "" {
		return errors.New("prepare needs -cinnabar and -out")
	}
	corpus, err := retail.LoadCorpus(*corpusPath)
	if err != nil {
		return err
	}
	if err := retail.CheckOraclePin(*cinnabar); err != nil {
		return err
	}
	batches, err := corpus.ReadFixtures(*cinnabar)
	if err != nil {
		return err
	}
	unknownID, err := retail.UnusedPacketID(retail.KnownPacket)
	if err != nil {
		return err
	}
	inputs, err := retail.BuildInputs(corpus, batches, unknownID)
	if err != nil {
		return err
	}
	crafting, err := retail.CraftingInputs()
	if err != nil {
		return err
	}
	large, err := retail.LargeInputs()
	if err != nil {
		return err
	}
	inputs = append(append(inputs, crafting...), large...)
	if err := retail.WriteInputs(*out, inputs); err != nil {
		return err
	}
	fmt.Printf("prepared %d inputs from %d fixtures; unknown-id uses packet id %d\n", len(inputs), len(corpus.Fixtures), unknownID)
	return nil
}

func decode(args []string) error {
	flags := flag.NewFlagSet("decode", flag.ExitOnError)
	out := flags.String("out", "", "directory prepare wrote")
	flags.Parse(args)
	inputs, err := retail.ReadInputs(*out)
	if err != nil {
		return err
	}
	var generated, oracle []retail.Verdict
	for _, input := range inputs {
		generated = append(generated, retail.DecodeGenerated(input)...)
		oracle = append(oracle, retail.DecodeGophertunnel(input)...)
	}
	if err := retail.WriteVerdicts(*out, retail.DecoderGo, generated); err != nil {
		return err
	}
	return retail.WriteVerdicts(*out, retail.DecoderGophertunnel, oracle)
}

func compare(args []string) error {
	flags := flag.NewFlagSet("compare", flag.ExitOnError)
	corpusPath := flags.String("corpus", "corpus.json", "retail corpus")
	exceptionsPath := flags.String("exceptions", "exceptions.json", "reviewed exceptions")
	goMod := flags.String("gomod", "go.mod", "this module's go.mod, which pins the oracle")
	out := flags.String("out", "", "directory holding expected.tsv and verdicts/")
	flags.Parse(args)
	corpus, err := retail.LoadCorpus(*corpusPath)
	if err != nil {
		return err
	}
	oracle, err := retail.GophertunnelReplacement(*goMod)
	if err != nil {
		return err
	}
	exceptions, err := retail.LoadExceptions(*exceptionsPath, corpus, oracle)
	if err != nil {
		return err
	}
	expected, err := retail.ReadExpected(*out)
	if err != nil {
		return err
	}
	verdicts := map[string][]retail.Verdict{}
	for _, decoder := range retail.Decoders {
		if verdicts[decoder], err = retail.ReadVerdicts(*out, decoder); err != nil {
			return err
		}
	}
	report := retail.Apply(retail.Compare(expected, verdicts), exceptions)
	text := formatReport(retail.Summarize(verdicts), report)
	fmt.Print(text)
	if err := os.WriteFile(filepath.Join(*out, "report.txt"), []byte(text), 0o644); err != nil {
		return err
	}
	if report.Failed() {
		return errFailed
	}
	return nil
}

func formatReport(summary string, report retail.Report) string {
	var b strings.Builder
	b.WriteString(summary)
	fmt.Fprintf(&b, "\n%d unexpected disagreements, %d accepted by exceptions, %d stale exceptions\n",
		len(report.Unexpected), len(report.Accepted), len(report.Stale))
	for _, d := range report.Unexpected {
		fmt.Fprintf(&b, "UNEXPECTED\t%s\t%s\t%s\t%d\t%s\t%s\n", d.Decoder, d.Case, d.Fixture, d.Entry, d.Kind, d.Detail)
	}
	for _, d := range report.Accepted {
		fmt.Fprintf(&b, "accepted\t%s\t%s\t%s\t%d\t%s\t%s\n", d.Decoder, d.Case, d.Fixture, d.Entry, d.Kind, d.Detail)
	}
	for _, stale := range report.Stale {
		e := stale.Exception
		fmt.Fprintf(&b, "STALE\t%s\t%s\t%s\t%s\t%s\n", e.Decoder, e.Case, stale.Fixture, e.Kind, e.Reason)
	}
	return b.String()
}
