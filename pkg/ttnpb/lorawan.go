// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttnpb

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/blang/semver"
	"github.com/vmihailenco/msgpack/v5"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

func init() {
	for i := range PHYVersion_name {
		PHYVersion_value[PHYVersion(i).String()] = i
	}
	PHYVersion_value["1.0"] = int32(PHY_V1_0)           // 1.0 is the official version number
	PHYVersion_value["1.0.2"] = int32(PHY_V1_0_2_REV_A) // Revisions were added from 1.0.2-b
	PHYVersion_value["1.1-a"] = int32(PHY_V1_1_REV_A)   // 1.1 is the official version number
	PHYVersion_value["1.1-b"] = int32(PHY_V1_1_REV_B)   // 1.1 is the official version number
}

// MarshalText implements encoding.TextMarshaler interface.
func (v MType) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *MType) UnmarshalText(b []byte) error {
	if i, ok := MType_value[string(b)]; ok {
		*v = MType(i)
		return nil
	}
	return errCouldNotParse("MType")(string(b))
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *MType) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("MType")(string(b)).WithCause(err)
	}
	*v = MType(i)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (v Major) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *Major) UnmarshalText(b []byte) error {
	if i, ok := Major_value[string(b)]; ok {
		*v = Major(i)
		return nil
	}
	return errCouldNotParse("Major")(string(b))
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *Major) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("Major")(string(b)).WithCause(err)
	}
	*v = Major(i)
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler interface.
func (v MACVersion) MarshalBinary() ([]byte, error) {
	if v > 255 {
		panic(fmt.Errorf("MACVersion enum exceeds 255"))
	}
	return []byte{byte(v)}, nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (v MACVersion) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// EncodeMsgpack implements msgpack.CustomEncoder interface.
func (v MACVersion) EncodeMsgpack(enc *msgpack.Encoder) error {
	if v > 255 {
		panic(fmt.Errorf("MACVersion enum exceeds 255"))
	}
	return enc.EncodeUint8(uint8(v))
}

// UnmarshalBinary implements encoding.BinaryMarshaler interface.
func (v *MACVersion) UnmarshalBinary(b []byte) error {
	if len(b) != 1 {
		return errCouldNotParse("MACVersion")(string(b))
	}
	*v = MACVersion(b[0])
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *MACVersion) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := MACVersion_value[s]; ok {
		*v = MACVersion(i)
		return nil
	}
	if !strings.HasPrefix(s, "MAC_") {
		if i, ok := MACVersion_value["MAC_"+s]; ok {
			*v = MACVersion(i)
			return nil
		}
	}
	return errCouldNotParse("MACVersion")(s)
}

// DecodeMsgpack implements msgpack.CustomDecoder interface.
func (v *MACVersion) DecodeMsgpack(dec *msgpack.Decoder) error {
	i, err := dec.DecodeInt32()
	if err != nil {
		return err
	}
	*v = MACVersion(i)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *MACVersion) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("MACVersion")(string(b)).WithCause(err)
	}
	*v = MACVersion(i)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (v PHYVersion) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *PHYVersion) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := PHYVersion_value[s]; ok {
		*v = PHYVersion(i)
		return nil
	}
	if !strings.HasPrefix(s, "PHY_") {
		if i, ok := PHYVersion_value["PHY_"+s]; ok {
			*v = PHYVersion(i)
			return nil
		}
	}
	return errCouldNotParse("PHYVersion")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *PHYVersion) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("PHYVersion")(string(b)).WithCause(err)
	}
	*v = PHYVersion(i)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (v DataRateIndex) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// MarshalJSON implements json.Marshaler interface.
func (v DataRateIndex) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *DataRateIndex) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := DataRateIndex_value[s]; ok {
		*v = DataRateIndex(i)
		return nil
	}
	if !strings.HasPrefix(s, "DATA_RATE_") {
		if i, ok := DataRateIndex_value["DATA_RATE_"+s]; ok {
			*v = DataRateIndex(i)
			return nil
		}
	}
	return errCouldNotParse("DataRateIndex")(string(b))
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *DataRateIndex) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("DataRateIndex")(string(b)).WithCause(err)
	}
	*v = DataRateIndex(i)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *RejoinType) UnmarshalText(b []byte) error {
	if i, ok := RejoinType_value[string(b)]; ok {
		*v = RejoinType(i)
		return nil
	}
	return errCouldNotParse("RejoinType")(string(b))
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *RejoinType) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("RejoinType")(string(b)).WithCause(err)
	}
	*v = RejoinType(i)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *CFListType) UnmarshalText(b []byte) error {
	if i, ok := CFListType_value[string(b)]; ok {
		*v = CFListType(i)
		return nil
	}
	return errCouldNotParse("CFListType")(string(b))
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *CFListType) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("CFListType")(string(b)).WithCause(err)
	}
	*v = CFListType(i)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (v Class) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *Class) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := Class_value[s]; ok {
		*v = Class(i)
		return nil
	}
	if !strings.HasPrefix(s, "CLASS_") {
		if i, ok := Class_value["CLASS_"+s]; ok {
			*v = Class(i)
			return nil
		}
	}
	return errCouldNotParse("Class")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *Class) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("Class")(string(b)).WithCause(err)
	}
	*v = Class(i)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (v TxSchedulePriority) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *TxSchedulePriority) UnmarshalText(b []byte) error {
	if i, ok := TxSchedulePriority_value[string(b)]; ok {
		*v = TxSchedulePriority(i)
		return nil
	}
	return errCouldNotParse("TxSchedulePriority")(string(b))
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *TxSchedulePriority) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("TxSchedulePriority")(string(b)).WithCause(err)
	}
	*v = TxSchedulePriority(i)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (v MACCommandIdentifier) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *MACCommandIdentifier) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := MACCommandIdentifier_value[s]; ok {
		*v = MACCommandIdentifier(i)
		return nil
	}
	if !strings.HasPrefix(s, "CID_") {
		if i, ok := MACCommandIdentifier_value["CID_"+s]; ok {
			*v = MACCommandIdentifier(i)
			return nil
		}
	}
	return errCouldNotParse("MACCommandIdentifier")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *MACCommandIdentifier) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("MACCommandIdentifier")(string(b)).WithCause(err)
	}
	*v = MACCommandIdentifier(i)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *AggregatedDutyCycle) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := AggregatedDutyCycle_value[s]; ok {
		*v = AggregatedDutyCycle(i)
		return nil
	}
	if !strings.HasPrefix(s, "DUTY_CYCLE_") {
		if i, ok := AggregatedDutyCycle_value["DUTY_CYCLE_"+s]; ok {
			*v = AggregatedDutyCycle(i)
			return nil
		}
	}
	return errCouldNotParse("AggregatedDutyCycle")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *AggregatedDutyCycle) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("AggregatedDutyCycle")(string(b)).WithCause(err)
	}
	*v = AggregatedDutyCycle(i)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (v PingSlotPeriod) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *PingSlotPeriod) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := PingSlotPeriod_value[s]; ok {
		*v = PingSlotPeriod(i)
		return nil
	}
	if !strings.HasPrefix(s, "PING_EVERY_") {
		if i, ok := PingSlotPeriod_value["PING_EVERY_"+s]; ok {
			*v = PingSlotPeriod(i)
			return nil
		}
	}
	return errCouldNotParse("PingSlotPeriod")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *PingSlotPeriod) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("PingSlotPeriod")(string(b)).WithCause(err)
	}
	*v = PingSlotPeriod(i)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *RejoinCountExponent) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := RejoinCountExponent_value[s]; ok {
		*v = RejoinCountExponent(i)
		return nil
	}
	if !strings.HasPrefix(s, "REJOIN_COUNT_") {
		if i, ok := RejoinCountExponent_value["REJOIN_COUNT_"+s]; ok {
			*v = RejoinCountExponent(i)
			return nil
		}
	}
	return errCouldNotParse("RejoinCountExponent")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *RejoinCountExponent) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("RejoinCountExponent")(string(b)).WithCause(err)
	}
	*v = RejoinCountExponent(i)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *RejoinTimeExponent) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := RejoinTimeExponent_value[s]; ok {
		*v = RejoinTimeExponent(i)
		return nil
	}
	if !strings.HasPrefix(s, "REJOIN_TIME_") {
		if i, ok := RejoinTimeExponent_value["REJOIN_TIME_"+s]; ok {
			*v = RejoinTimeExponent(i)
			return nil
		}
	}
	return errCouldNotParse("RejoinTimeExponent")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *RejoinTimeExponent) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("RejoinTimeExponent")(string(b)).WithCause(err)
	}
	*v = RejoinTimeExponent(i)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *RejoinPeriodExponent) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := RejoinPeriodExponent_value[s]; ok {
		*v = RejoinPeriodExponent(i)
		return nil
	}
	if !strings.HasPrefix(s, "REJOIN_PERIOD_") {
		if i, ok := RejoinPeriodExponent_value["REJOIN_PERIOD_"+s]; ok {
			*v = RejoinPeriodExponent(i)
			return nil
		}
	}
	return errCouldNotParse("RejoinPeriodExponent")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *RejoinPeriodExponent) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("RejoinPeriodExponent")(string(b)).WithCause(err)
	}
	*v = RejoinPeriodExponent(i)
	return nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (v DeviceEIRP) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *DeviceEIRP) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := DeviceEIRP_value[s]; ok {
		*v = DeviceEIRP(i)
		return nil
	}
	if !strings.HasPrefix(s, "DEVICE_EIRP_") {
		if i, ok := DeviceEIRP_value["DEVICE_EIRP_"+s]; ok {
			*v = DeviceEIRP(i)
			return nil
		}
	}
	return errCouldNotParse("DeviceEIRP")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *DeviceEIRP) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("DeviceEIRP")(string(b)).WithCause(err)
	}
	*v = DeviceEIRP(i)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *ADRAckLimitExponent) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := ADRAckLimitExponent_value[s]; ok {
		*v = ADRAckLimitExponent(i)
		return nil
	}
	if !strings.HasPrefix(s, "ADR_ACK_LIMIT_") {
		if i, ok := ADRAckLimitExponent_value["ADR_ACK_LIMIT_"+s]; ok {
			*v = ADRAckLimitExponent(i)
			return nil
		}
	}
	return errCouldNotParse("ADRAckLimitExponent")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *ADRAckLimitExponent) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("ADRAckLimitExponent")(string(b)).WithCause(err)
	}
	*v = ADRAckLimitExponent(i)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *ADRAckDelayExponent) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := ADRAckDelayExponent_value[s]; ok {
		*v = ADRAckDelayExponent(i)
		return nil
	}
	if !strings.HasPrefix(s, "ADR_ACK_DELAY_") {
		if i, ok := ADRAckDelayExponent_value["ADR_ACK_DELAY_"+s]; ok {
			*v = ADRAckDelayExponent(i)
			return nil
		}
	}
	return errCouldNotParse("ADRAckDelayExponent")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *ADRAckDelayExponent) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("ADRAckDelayExponent")(string(b)).WithCause(err)
	}
	*v = ADRAckDelayExponent(i)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *RxDelay) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := RxDelay_value[s]; ok {
		*v = RxDelay(i)
		return nil
	}
	if !strings.HasPrefix(s, "RX_DELAY_") {
		if i, ok := RxDelay_value["RX_DELAY_"+s]; ok {
			*v = RxDelay(i)
			return nil
		}
	}
	return errCouldNotParse("RxDelay")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *RxDelay) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("RxDelay")(string(b)).WithCause(err)
	}
	*v = RxDelay(i)
	return nil
}

// Duration returns v as time.Duration.
func (v RxDelay) Duration() time.Duration {
	switch v {
	case RX_DELAY_0, RX_DELAY_1:
		return time.Second
	default:
		return time.Duration(v) * time.Second
	}
}

// Validate reports whether v represents a valid RxDelay.
func (v RxDelay) Validate() error {
	if v < 0 || v >= RxDelay(len(RxDelay_name)) {
		return errExpectedBetween("RxDelay", 0, len(RxDelay_name)-1)(v)
	}
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (v *Minor) UnmarshalText(b []byte) error {
	s := string(b)
	if i, ok := Minor_value[s]; ok {
		*v = Minor(i)
		return nil
	}
	if !strings.HasPrefix(s, "MINOR_") {
		if i, ok := Minor_value["MINOR_"+s]; ok {
			*v = Minor(i)
			return nil
		}
	}
	return errCouldNotParse("Minor")(s)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (v *Minor) UnmarshalJSON(b []byte) error {
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		return v.UnmarshalText(b[1 : len(b)-1])
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return errCouldNotParse("Minor")(string(b)).WithCause(err)
	}
	*v = Minor(i)
	return nil
}

var errParsingSemanticVersion = unexpectedValue(
	errors.DefineInvalidArgument("parsing_semantic_version", "could not parse semantic version", valueKey),
)

// Validate reports whether v represents a valid MACVersion.
func (v MACVersion) Validate() error {
	if v < 1 || v >= MACVersion(len(MACVersion_name)) {
		return errExpectedBetween("MACVersion", 1, len(MACVersion_name)-1)(v)
	}

	_, err := semver.Parse(v.String())
	if err != nil {
		return errParsingSemanticVersion(v.String()).WithCause(err)
	}
	return nil
}

// String implements fmt.Stringer.
func (v MACVersion) String() string {
	switch v {
	case MAC_V1_0:
		return "1.0.0"
	case MAC_V1_0_1:
		return "1.0.1"
	case MAC_V1_0_2:
		return "1.0.2"
	case MAC_V1_0_3:
		return "1.0.3"
	case MAC_V1_0_4:
		return "1.0.4"
	case MAC_V1_1:
		return "1.1.0"
	}
	return "unknown"
}

func init() {
	for i := range MACVersion_name {
		MACVersion_value[MACVersion(i).String()] = i
	}
	MACVersion_value["1.0"] = int32(MAC_V1_0) // 1.0 is the official version number
	MACVersion_value["1.1"] = int32(MAC_V1_1) // 1.1 is the official version number
}

// Compare compares MACVersions v to o:
// -1 == v is less than o
// 0 == v is equal to o
// 1 == v is greater than o
// Compare panics, if v.Validate() returns non-nil error.
func (v MACVersion) Compare(o MACVersion) int {
	return semver.MustParse(v.String()).Compare(
		semver.MustParse(o.String()),
	)
}

// EncryptFOpts reports whether v requires MAC commands in FOpts to be encrypted.
// EncryptFOpts panics, if v.Validate() returns non-nil error.
func (v MACVersion) EncryptFOpts() bool {
	return v.Compare(MAC_V1_1) >= 0
}

// HasMaxFCntGap reports whether v defines a MaxFCntGap.
// HasMaxFCntGap panics, if v.Validate() returns non-nil error.
func (v MACVersion) HasMaxFCntGap() bool {
	return v.Compare(MAC_V1_0_4) < 0
}

// HasNoChangeTXPowerIndex reports whether v defines a no-change TxPowerIndex value.
// HasNoChangeTXPowerIndex panics, if v.Validate() returns non-nil error.
func (v MACVersion) HasNoChangeTXPowerIndex() bool {
	return v.Compare(MAC_V1_0_4) >= 0
}

// HasNoChangeDataRateIndex reports whether v defines a no-change DataRateIndex value.
// HasNoChangeDataRateIndex panics, if v.Validate() returns non-nil error.
func (v MACVersion) HasNoChangeDataRateIndex() bool {
	return v.Compare(MAC_V1_0_4) >= 0
}

// IgnoreUplinksExceedingLengthLimit reports whether v requires Network Server to
// silently drop uplinks exceeding selected data rate payload length limits.
// IgnoreUplinksExceedingLengthLimit panics, if v.Validate() returns non-nil error.
func (v MACVersion) IgnoreUplinksExceedingLengthLimit() bool {
	return v.Compare(MAC_V1_0_4) >= 0 && v.Compare(MAC_V1_1) < 0
}

// IncrementDevNonce reports whether v defines DevNonce as an incrementing counter.
// IncrementDevNonce panics, if v.Validate() returns non-nil error.
func (v MACVersion) IncrementDevNonce() bool {
	return v.Compare(MAC_V1_0_4) >= 0
}

// UseNwkKey reports whether v uses a root NwkKey.
// UseNwkKey panics, if v.Validate() returns non-nil error.
func (v MACVersion) UseNwkKey() bool {
	return v.Compare(MAC_V1_1) >= 0
}

// UseLegacyMIC reports whether v uses legacy MIC computation algorithm.
// UseLegacyMIC panics, if v.Validate() returns non-nil error.
func (v MACVersion) UseLegacyMIC() bool {
	return v.Compare(MAC_V1_1) < 0
}

// RequireDevEUIForABP reports whether v requires ABP devices to have a DevEUI associated.
// RequireDevEUIForABP panics, if v.Validate() returns non-nil error.
func (v MACVersion) RequireDevEUIForABP() bool {
	return v.Compare(MAC_V1_0_4) >= 0 && v.Compare(MAC_V1_1) < 0
}

// Validate reports whether v represents a valid PHYVersion.
func (v PHYVersion) Validate() error {
	if v < 1 || v >= PHYVersion(len(PHYVersion_name)) {
		return errExpectedBetween("PHYVersion", 1, len(PHYVersion_name)-1)(v)
	}

	_, err := semver.Parse(v.String())
	if err != nil {
		return errParsingSemanticVersion(v.String()).WithCause(err)
	}
	return nil
}

// String implements fmt.Stringer.
func (v PHYVersion) String() string {
	switch v {
	case PHY_V1_0:
		return "1.0.0"
	case PHY_V1_0_1:
		return "1.0.1"
	case PHY_V1_0_2_REV_A:
		return "1.0.2-a"
	case PHY_V1_0_2_REV_B:
		return "1.0.2-b"
	case PHY_V1_0_3_REV_A:
		return "1.0.3-a"
	case PHY_V1_1_REV_A:
		return "1.1.0-a"
	case PHY_V1_1_REV_B:
		return "1.1.0-b"
	}
	return "unknown"
}

// Compare compares PHYVersions v to o:
// -1 == v is less than o
// 0 == v is equal to o
// 1 == v is greater than o
// Compare panics, if v.Validate() returns non-nil error.
func (v PHYVersion) Compare(o PHYVersion) int {
	return semver.MustParse(v.String()).Compare(
		semver.MustParse(o.String()),
	)
}

// String implements fmt.Stringer.
func (v DataRateIndex) String() string {
	return strconv.Itoa(int(v))
}

// String implements fmt.Stringer.
func (v RxDelay) String() string {
	return strconv.Itoa(int(v))
}

func (v LoRaDataRate) DataRate() DataRate {
	return DataRate{
		Modulation: &DataRate_LoRa{
			LoRa: &v,
		},
	}
}

func (v FSKDataRate) DataRate() DataRate {
	return DataRate{
		Modulation: &DataRate_FSK{
			FSK: &v,
		},
	}
}

// FieldIsZero returns whether path p is zero.
func (v *CFList) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "ch_masks":
		return v.ChMasks == nil
	case "freq":
		return v.Freq == nil
	case "type":
		return v.Type == 0
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}

// FieldIsZero returns whether path p is zero.
func (v *DLSettings) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "opt_neg":
		return !v.OptNeg
	case "rx1_dr_offset":
		return v.Rx1DROffset == 0
	case "rx2_dr":
		return v.Rx2DR == 0
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}

// FieldIsZero returns whether path p is zero.
func (v *MHDR) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "m_type":
		return v.MType == 0
	case "major":
		return v.Major == 0
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}

// FieldIsZero returns whether path p is zero.
func (v *JoinAcceptPayload) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "cf_list":
		return v.CFList == nil
	case "cf_list.ch_masks":
		return v.CFList.FieldIsZero("ch_masks")
	case "cf_list.freq":
		return v.CFList.FieldIsZero("freq")
	case "cf_list.type":
		return v.CFList.FieldIsZero("type")
	case "dev_addr":
		return v.DevAddr == types.DevAddr{}
	case "dl_settings":
		return v.DLSettings == DLSettings{}
	case "dl_settings.opt_neg":
		return v.DLSettings.FieldIsZero("opt_neg")
	case "dl_settings.rx1_dr_offset":
		return v.DLSettings.FieldIsZero("rx1_dr_offset")
	case "dl_settings.rx2_dr":
		return v.DLSettings.FieldIsZero("rx2_dr")
	case "encrypted":
		return v.Encrypted == nil
	case "join_nonce":
		return v.JoinNonce == types.JoinNonce{}
	case "net_id":
		return v.NetID == types.NetID{}
	case "rx_delay":
		return v.RxDelay == 0
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}

// FieldIsZero returns whether path p is zero.
func (v *JoinRequestPayload) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "dev_eui":
		return v.DevEUI == types.EUI64{}
	case "dev_nonce":
		return v.DevNonce == types.DevNonce{}
	case "join_eui":
		return v.JoinEUI == types.EUI64{}
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}

// FieldIsZero returns whether path p is zero.
func (v *FCtrl) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "ack":
		return !v.Ack
	case "adr":
		return !v.ADR
	case "adr_ack_req":
		return !v.ADRAckReq
	case "class_b":
		return !v.ClassB
	case "f_pending":
		return !v.FPending
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}

// FieldIsZero returns whether path p is zero.
func (v *FHDR) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "dev_addr":
		return v.DevAddr == types.DevAddr{}
	case "f_cnt":
		return v.FCnt == 0
	case "f_ctrl":
		return v.FCtrl == FCtrl{}
	case "f_ctrl.ack":
		return v.FCtrl.FieldIsZero("ack")
	case "f_ctrl.adr":
		return v.FCtrl.FieldIsZero("adr")
	case "f_ctrl.adr_ack_req":
		return v.FCtrl.FieldIsZero("adr_ack_req")
	case "f_ctrl.class_b":
		return v.FCtrl.FieldIsZero("class_b")
	case "f_ctrl.f_pending":
		return v.FCtrl.FieldIsZero("f_pending")
	case "f_opts":
		return v.FOpts == nil
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}

// FieldIsZero returns whether path p is zero.
func (v *MACPayload) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "decoded_payload":
		return v.DecodedPayload == nil
	case "f_hdr":
		return fieldsAreZero(&v.FHDR, FHDRFieldPathsTopLevel...)
	case "f_hdr.dev_addr":
		return v.FHDR.FieldIsZero("dev_addr")
	case "f_hdr.f_cnt":
		return v.FHDR.FieldIsZero("f_cnt")
	case "f_hdr.f_ctrl":
		return v.FHDR.FieldIsZero("f_ctrl")
	case "f_hdr.f_ctrl.ack":
		return v.FHDR.FieldIsZero("f_ctrl.ack")
	case "f_hdr.f_ctrl.adr":
		return v.FHDR.FieldIsZero("f_ctrl.adr")
	case "f_hdr.f_ctrl.adr_ack_req":
		return v.FHDR.FieldIsZero("f_ctrl.adr_ack_req")
	case "f_hdr.f_ctrl.class_b":
		return v.FHDR.FieldIsZero("f_ctrl.class_b")
	case "f_hdr.f_ctrl.f_pending":
		return v.FHDR.FieldIsZero("f_ctrl.f_pending")
	case "f_hdr.f_opts":
		return v.FHDR.FieldIsZero("f_opts")
	case "f_port":
		return v.FPort == 0
	case "frm_payload":
		return v.FRMPayload == nil
	case "full_f_cnt":
		return v.FullFCnt == 0
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}

// FieldIsZero returns whether path p is zero.
func (v *RejoinRequestPayload) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "dev_eui":
		return v.DevEUI == types.EUI64{}
	case "join_eui":
		return v.JoinEUI == types.EUI64{}
	case "net_id":
		return v.NetID == types.NetID{}
	case "rejoin_cnt":
		return v.RejoinCnt == 0
	case "rejoin_type":
		return v.RejoinType == 0
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}

// FieldIsZero returns whether path p is zero.
func (v *Message) FieldIsZero(p string) bool {
	if v == nil {
		return true
	}
	switch p {
	case "Payload":
		return v.Payload == nil
	case "Payload.join_accept_payload":
		return v.GetJoinAcceptPayload() == nil
	case "Payload.join_accept_payload.cf_list":
		return v.GetJoinAcceptPayload().FieldIsZero("cf_list")
	case "Payload.join_accept_payload.cf_list.ch_masks":
		return v.GetJoinAcceptPayload().FieldIsZero("cf_list.ch_masks")
	case "Payload.join_accept_payload.cf_list.freq":
		return v.GetJoinAcceptPayload().FieldIsZero("cf_list.freq")
	case "Payload.join_accept_payload.cf_list.type":
		return v.GetJoinAcceptPayload().FieldIsZero("cf_list.type")
	case "Payload.join_accept_payload.dev_addr":
		return v.GetJoinAcceptPayload().FieldIsZero("dev_addr")
	case "Payload.join_accept_payload.dl_settings":
		return v.GetJoinAcceptPayload().FieldIsZero("dl_settings")
	case "Payload.join_accept_payload.dl_settings.opt_neg":
		return v.GetJoinAcceptPayload().FieldIsZero("dl_settings.opt_neg")
	case "Payload.join_accept_payload.dl_settings.rx1_dr_offset":
		return v.GetJoinAcceptPayload().FieldIsZero("dl_settings.rx1_dr_offset")
	case "Payload.join_accept_payload.dl_settings.rx2_dr":
		return v.GetJoinAcceptPayload().FieldIsZero("dl_settings.rx2_dr")
	case "Payload.join_accept_payload.encrypted":
		return v.GetJoinAcceptPayload().FieldIsZero("encrypted")
	case "Payload.join_accept_payload.join_nonce":
		return v.GetJoinAcceptPayload().FieldIsZero("join_nonce")
	case "Payload.join_accept_payload.net_id":
		return v.GetJoinAcceptPayload().FieldIsZero("net_id")
	case "Payload.join_accept_payload.rx_delay":
		return v.GetJoinAcceptPayload().FieldIsZero("rx_delay")
	case "Payload.join_request_payload":
		return v.GetJoinRequestPayload() == nil
	case "Payload.join_request_payload.dev_eui":
		return v.GetJoinRequestPayload().FieldIsZero("dev_eui")
	case "Payload.join_request_payload.dev_nonce":
		return v.GetJoinRequestPayload().FieldIsZero("dev_nonce")
	case "Payload.join_request_payload.join_eui":
		return v.GetJoinRequestPayload().FieldIsZero("join_eui")
	case "Payload.mac_payload":
		return v.GetMACPayload() == nil
	case "Payload.mac_payload.decoded_payload":
		return v.GetMACPayload().FieldIsZero("decoded_payload")
	case "Payload.mac_payload.f_hdr":
		return v.GetMACPayload().FieldIsZero("f_hdr")
	case "Payload.mac_payload.f_hdr.dev_addr":
		return v.GetMACPayload().FieldIsZero("f_hdr.dev_addr")
	case "Payload.mac_payload.f_hdr.f_cnt":
		return v.GetMACPayload().FieldIsZero("f_hdr.f_cnt")
	case "Payload.mac_payload.f_hdr.f_ctrl":
		return v.GetMACPayload().FieldIsZero("f_hdr.f_ctrl")
	case "Payload.mac_payload.f_hdr.f_ctrl.ack":
		return v.GetMACPayload().FieldIsZero("f_hdr.f_ctrl.ack")
	case "Payload.mac_payload.f_hdr.f_ctrl.adr":
		return v.GetMACPayload().FieldIsZero("f_hdr.f_ctrl.adr")
	case "Payload.mac_payload.f_hdr.f_ctrl.adr_ack_req":
		return v.GetMACPayload().FieldIsZero("f_hdr.f_ctrl.adr_ack_req")
	case "Payload.mac_payload.f_hdr.f_ctrl.class_b":
		return v.GetMACPayload().FieldIsZero("f_hdr.f_ctrl.class_b")
	case "Payload.mac_payload.f_hdr.f_ctrl.f_pending":
		return v.GetMACPayload().FieldIsZero("f_hdr.f_ctrl.f_pending")
	case "Payload.mac_payload.f_hdr.f_opts":
		return v.GetMACPayload().FieldIsZero("f_hdr.f_opts")
	case "Payload.mac_payload.f_port":
		return v.GetMACPayload().FieldIsZero("f_port")
	case "Payload.mac_payload.frm_payload":
		return v.GetMACPayload().FieldIsZero("frm_payload")
	case "Payload.mac_payload.full_f_cnt":
		return v.GetMACPayload().FieldIsZero("full_f_cnt")
	case "Payload.rejoin_request_payload":
		return v.GetRejoinRequestPayload() == nil
	case "Payload.rejoin_request_payload.dev_eui":
		return v.GetRejoinRequestPayload().FieldIsZero("dev_eui")
	case "Payload.rejoin_request_payload.join_eui":
		return v.GetRejoinRequestPayload().FieldIsZero("join_eui")
	case "Payload.rejoin_request_payload.net_id":
		return v.GetRejoinRequestPayload().FieldIsZero("net_id")
	case "Payload.rejoin_request_payload.rejoin_cnt":
		return v.GetRejoinRequestPayload().FieldIsZero("rejoin_cnt")
	case "Payload.rejoin_request_payload.rejoin_type":
		return v.GetRejoinRequestPayload().FieldIsZero("rejoin_type")
	case "m_hdr":
		return v.MHDR == MHDR{}
	case "m_hdr.m_type":
		return v.MHDR.FieldIsZero("m_type")
	case "m_hdr.major":
		return v.MHDR.FieldIsZero("major")
	case "mic":
		return v.MIC == nil
	}
	panic(fmt.Sprintf("unknown path '%s'", p))
}
