// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackNetResult uint8

const (
	ItemStackNetResultSuccess                                          ItemStackNetResult = 0
	ItemStackNetResultError                                            ItemStackNetResult = 1
	ItemStackNetResultInvalidRequestActionType                         ItemStackNetResult = 2
	ItemStackNetResultActionRequestNotAllowed                          ItemStackNetResult = 3
	ItemStackNetResultScreenHandlerEndRequestFailed                    ItemStackNetResult = 4
	ItemStackNetResultItemRequestActionHandlerCommitFailed             ItemStackNetResult = 5
	ItemStackNetResultInvalidRequestCraftActionType                    ItemStackNetResult = 6
	ItemStackNetResultInvalidCraftRequest                              ItemStackNetResult = 7
	ItemStackNetResultInvalidCraftRequestScreen                        ItemStackNetResult = 8
	ItemStackNetResultInvalidCraftResult                               ItemStackNetResult = 9
	ItemStackNetResultInvalidCraftResultIndex                          ItemStackNetResult = 10
	ItemStackNetResultInvalidCraftResultItem                           ItemStackNetResult = 11
	ItemStackNetResultInvalidItemNetId                                 ItemStackNetResult = 12
	ItemStackNetResultMissingCreatedOutputContainer                    ItemStackNetResult = 13
	ItemStackNetResultFailedToSetCreatedItemOutputSlot                 ItemStackNetResult = 14
	ItemStackNetResultRequestAlreadyInProgress                         ItemStackNetResult = 15
	ItemStackNetResultFailedToInitSparseContainer                      ItemStackNetResult = 16
	ItemStackNetResultResultTransferFailed                             ItemStackNetResult = 17
	ItemStackNetResultExpectedItemSlotNotFullyConsumed                 ItemStackNetResult = 18
	ItemStackNetResultExpectedAnywhereItemNotFullyConsumed             ItemStackNetResult = 19
	ItemStackNetResultItemAlreadyConsumedFromSlot                      ItemStackNetResult = 20
	ItemStackNetResultConsumedTooMuchFromSlot                          ItemStackNetResult = 21
	ItemStackNetResultMismatchSlotExpectedConsumedItem                 ItemStackNetResult = 22
	ItemStackNetResultMismatchSlotExpectedConsumedItemNetIdVariant     ItemStackNetResult = 23
	ItemStackNetResultFailedToMatchExpectedSlotConsumedItem            ItemStackNetResult = 24
	ItemStackNetResultFailedToMatchExpectedAllowedAnywhereConsumedItem ItemStackNetResult = 25
	ItemStackNetResultConsumedItemOutOfAllowedSlotRange                ItemStackNetResult = 26
	ItemStackNetResultConsumedItemNotAllowed                           ItemStackNetResult = 27
	ItemStackNetResultPlayerNotInCreativeMode                          ItemStackNetResult = 28
	ItemStackNetResultInvalidExperimentalRecipeRequest                 ItemStackNetResult = 29
	ItemStackNetResultFailedToCraftCreative                            ItemStackNetResult = 30
	ItemStackNetResultFailedToGetLevelRecipe                           ItemStackNetResult = 31
	ItemStackNetResultFailedToFindRecipeByNetId                        ItemStackNetResult = 32
	ItemStackNetResultMismatchedCraftingSize                           ItemStackNetResult = 33
	ItemStackNetResultMissingInputSparseContainer                      ItemStackNetResult = 34
	ItemStackNetResultMismatchedRecipeForInputGridItems                ItemStackNetResult = 35
	ItemStackNetResultEmptyCraftResults                                ItemStackNetResult = 36
	ItemStackNetResultFailedToEnchant                                  ItemStackNetResult = 37
	ItemStackNetResultMissingInputItem                                 ItemStackNetResult = 38
	ItemStackNetResultInsufficientPlayerLevelToEnchant                 ItemStackNetResult = 39
	ItemStackNetResultMissingMaterialItem                              ItemStackNetResult = 40
	ItemStackNetResultMissingActor                                     ItemStackNetResult = 41
	ItemStackNetResultUnknownPrimaryEffect                             ItemStackNetResult = 42
	ItemStackNetResultPrimaryEffectOutOfRange                          ItemStackNetResult = 43
	ItemStackNetResultPrimaryEffectUnavailable                         ItemStackNetResult = 44
	ItemStackNetResultSecondaryEffectOutOfRange                        ItemStackNetResult = 45
	ItemStackNetResultSecondaryEffectUnavailable                       ItemStackNetResult = 46
	ItemStackNetResultDstContainerEqualToCreatedOutputContainer        ItemStackNetResult = 47
	ItemStackNetResultDstContainerAndSlotEqualToSrcContainerAndSlot    ItemStackNetResult = 48
	ItemStackNetResultFailedToValidateSrcSlot                          ItemStackNetResult = 49
	ItemStackNetResultFailedToValidateDstSlot                          ItemStackNetResult = 50
	ItemStackNetResultInvalidAdjustedAmount                            ItemStackNetResult = 51
	ItemStackNetResultInvalidItemSetType                               ItemStackNetResult = 52
	ItemStackNetResultInvalidTransferAmount                            ItemStackNetResult = 53
	ItemStackNetResultCannotSwapItem                                   ItemStackNetResult = 54
	ItemStackNetResultCannotPlaceItem                                  ItemStackNetResult = 55
	ItemStackNetResultUnhandledItemSetType                             ItemStackNetResult = 56
	ItemStackNetResultInvalidRemovedAmount                             ItemStackNetResult = 57
	ItemStackNetResultInvalidRegion                                    ItemStackNetResult = 58
	ItemStackNetResultCannotDropItem                                   ItemStackNetResult = 59
	ItemStackNetResultCannotDestroyItem                                ItemStackNetResult = 60
	ItemStackNetResultInvalidSourceContainer                           ItemStackNetResult = 61
	ItemStackNetResultItemNotConsumed                                  ItemStackNetResult = 62
	ItemStackNetResultInvalidNumCrafts                                 ItemStackNetResult = 63
	ItemStackNetResultInvalidCraftResultStackSize                      ItemStackNetResult = 64
	ItemStackNetResultCannotRemoveItem                                 ItemStackNetResult = 65
	ItemStackNetResultCannotConsumeItem                                ItemStackNetResult = 66
	ItemStackNetResultScreenStackError                                 ItemStackNetResult = 67
)
