use bedrock_protocol_1_26_40::{
    packets::ResourcePackClientResponse,
    types::ResourcePackClientResponseData,
    wire::{Decode, Encode},
};

fn assert_wire(response: ResourcePackClientResponseData, expected: &[u8]) {
    let packet = ResourcePackClientResponse { response };
    let encoded = packet.encode_to_vec();
    assert_eq!(encoded, expected);
    assert_eq!(
        ResourcePackClientResponse::decode_exact(expected),
        Ok(packet)
    );
}

#[test]
fn resource_pack_client_response_uses_zero_based_discriminants() {
    assert_wire(
        ResourcePackClientResponseData::Cancel {
            response_type: "cancel".into(),
        },
        &[0, 6, b'c', b'a', b'n', b'c', b'e', b'l'],
    );
    assert_wire(
        ResourcePackClientResponseData::Downloading {
            response_type: "downloading".into(),
            downloading_packs: vec!["a_1.0.0".into()],
        },
        &[
            1, 11, b'd', b'o', b'w', b'n', b'l', b'o', b'a', b'd', b'i', b'n', b'g', 1, 7, b'a',
            b'_', b'1', b'.', b'0', b'.', b'0',
        ],
    );
    assert_wire(
        ResourcePackClientResponseData::DownloadingFinished {
            response_type: "downloadingfinished".into(),
        },
        &[
            2, 19, b'd', b'o', b'w', b'n', b'l', b'o', b'a', b'd', b'i', b'n', b'g', b'f', b'i',
            b'n', b'i', b's', b'h', b'e', b'd',
        ],
    );
    assert_wire(
        ResourcePackClientResponseData::ResourcePackStackFinished {
            response_type: "resourcepackstackfinished".into(),
        },
        &[
            3, 25, b'r', b'e', b's', b'o', b'u', b'r', b'c', b'e', b'p', b'a', b'c', b'k', b's',
            b't', b'a', b'c', b'k', b'f', b'i', b'n', b'i', b's', b'h', b'e', b'd',
        ],
    );
}
