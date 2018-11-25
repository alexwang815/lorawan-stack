// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

var _RxMetadataFieldPaths = [...]string{
	"advanced",
	"antenna_index",
	"channel_rssi",
	"downlink_path_constraint",
	"encrypted_fine_timestamp",
	"encrypted_fine_timestamp_key_id",
	"fine_timestamp",
	"frequency_offset",
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
	"location",
	"location.accuracy",
	"location.altitude",
	"location.latitude",
	"location.longitude",
	"location.source",
	"rssi",
	"rssi_standard_deviation",
	"snr",
	"time",
	"timestamp",
}

func (*RxMetadata) FieldMaskPaths() []string {
	ret := make([]string, len(_RxMetadataFieldPaths))
	copy(ret, _RxMetadataFieldPaths[:])
	return ret
}

func (dst *RxMetadata) SetFields(src *RxMetadata, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "advanced":
			dst.Advanced = src.Advanced
		case "antenna_index":
			dst.AntennaIndex = src.AntennaIndex
		case "channel_rssi":
			dst.ChannelRSSI = src.ChannelRSSI
		case "downlink_path_constraint":
			dst.DownlinkPathConstraint = src.DownlinkPathConstraint
		case "encrypted_fine_timestamp":
			dst.EncryptedFineTimestamp = src.EncryptedFineTimestamp
		case "encrypted_fine_timestamp_key_id":
			dst.EncryptedFineTimestampKeyID = src.EncryptedFineTimestampKeyID
		case "fine_timestamp":
			dst.FineTimestamp = src.FineTimestamp
		case "frequency_offset":
			dst.FrequencyOffset = src.FrequencyOffset
		case "gateway_ids":
			dst.GatewayIdentifiers = src.GatewayIdentifiers
		case "gateway_ids.eui":
			dst.GatewayIdentifiers.SetFields(&src.GatewayIdentifiers, _pathsWithoutPrefix("gateway_ids", paths)...)
		case "gateway_ids.gateway_id":
			dst.GatewayIdentifiers.SetFields(&src.GatewayIdentifiers, _pathsWithoutPrefix("gateway_ids", paths)...)
		case "location":
			dst.Location = src.Location
		case "location.accuracy":
			if dst.Location == nil {
				dst.Location = &Location{}
			}
			dst.Location.SetFields(src.Location, _pathsWithoutPrefix("location", paths)...)
		case "location.altitude":
			if dst.Location == nil {
				dst.Location = &Location{}
			}
			dst.Location.SetFields(src.Location, _pathsWithoutPrefix("location", paths)...)
		case "location.latitude":
			if dst.Location == nil {
				dst.Location = &Location{}
			}
			dst.Location.SetFields(src.Location, _pathsWithoutPrefix("location", paths)...)
		case "location.longitude":
			if dst.Location == nil {
				dst.Location = &Location{}
			}
			dst.Location.SetFields(src.Location, _pathsWithoutPrefix("location", paths)...)
		case "location.source":
			if dst.Location == nil {
				dst.Location = &Location{}
			}
			dst.Location.SetFields(src.Location, _pathsWithoutPrefix("location", paths)...)
		case "rssi":
			dst.RSSI = src.RSSI
		case "rssi_standard_deviation":
			dst.RSSIStandardDeviation = src.RSSIStandardDeviation
		case "snr":
			dst.SNR = src.SNR
		case "time":
			dst.Time = src.Time
		case "timestamp":
			dst.Timestamp = src.Timestamp
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _TxMetadataFieldPaths = [...]string{
	"advanced",
	"antenna_index",
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
	"time",
	"timestamp",
}

func (*TxMetadata) FieldMaskPaths() []string {
	ret := make([]string, len(_TxMetadataFieldPaths))
	copy(ret, _TxMetadataFieldPaths[:])
	return ret
}

func (dst *TxMetadata) SetFields(src *TxMetadata, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "advanced":
			dst.Advanced = src.Advanced
		case "antenna_index":
			dst.AntennaIndex = src.AntennaIndex
		case "gateway_ids":
			dst.GatewayIdentifiers = src.GatewayIdentifiers
		case "gateway_ids.eui":
			dst.GatewayIdentifiers.SetFields(&src.GatewayIdentifiers, _pathsWithoutPrefix("gateway_ids", paths)...)
		case "gateway_ids.gateway_id":
			dst.GatewayIdentifiers.SetFields(&src.GatewayIdentifiers, _pathsWithoutPrefix("gateway_ids", paths)...)
		case "time":
			dst.Time = src.Time
		case "timestamp":
			dst.Timestamp = src.Timestamp
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _LocationFieldPaths = [...]string{
	"accuracy",
	"altitude",
	"latitude",
	"longitude",
	"source",
}

func (*Location) FieldMaskPaths() []string {
	ret := make([]string, len(_LocationFieldPaths))
	copy(ret, _LocationFieldPaths[:])
	return ret
}

func (dst *Location) SetFields(src *Location, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "accuracy":
			dst.Accuracy = src.Accuracy
		case "altitude":
			dst.Altitude = src.Altitude
		case "latitude":
			dst.Latitude = src.Latitude
		case "longitude":
			dst.Longitude = src.Longitude
		case "source":
			dst.Source = src.Source
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}