// Code generated by ogen, DO NOT EDIT.

package v1

import (
	"github.com/go-faster/errors"
)

type CloudCtrlAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *CloudCtrlAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *CloudCtrlAuth) SetToken(val string) {
	s.Token = val
}

type ConfigureProcessConfigurationBadRequest Error

func (*ConfigureProcessConfigurationBadRequest) configureProcessConfigurationRes() {}

type ConfigureProcessConfigurationInternalServerError Error

func (*ConfigureProcessConfigurationInternalServerError) configureProcessConfigurationRes() {}

type ConfigureProcessConfigurationNotFound Error

func (*ConfigureProcessConfigurationNotFound) configureProcessConfigurationRes() {}

type ConfigureProcessConfigurationOK struct {
	ProcessConfiguration ProcessConfiguration `json:"CommonServiceItem"`
}

// GetProcess returns the value of Process.
func (s *ConfigureProcessConfigurationOK) GetProcessConfiguration() ProcessConfiguration {
	return s.ProcessConfiguration
}

// SetProcess sets the value of Process.
func (s *ConfigureProcessConfigurationOK) SetProcessConfiguration(val ProcessConfiguration) {
	s.ProcessConfiguration = val
}

func (*ConfigureProcessConfigurationOK) configureProcessConfigurationRes() {}

// Ref: #/components/schemas/ConfigureProcessConfigurationRequest
type ConfigureProcessConfigurationRequest struct {
	ProcessConfiguration ProcessConfigurationRequestSettings `json:"CommonServiceItem"`
}

// GetName returns the value of Name.
func (s *ConfigureProcessConfigurationRequest) GetName() string {
	return s.ProcessConfiguration.Name
}

// GetDestination returns the value of Destination.
func (s *ConfigureProcessConfigurationRequest) GetDestination() ConfigureProcessConfigurationRequestDestination {
	return ConfigureProcessConfigurationRequestDestination(s.ProcessConfiguration.Settings.Destination)
}

// GetParameters returns the value of Parameters.
func (s *ConfigureProcessConfigurationRequest) GetParameters() string {
	return s.ProcessConfiguration.Settings.Parameters
}

// SetName sets the value of Name.
func (s *ConfigureProcessConfigurationRequest) SetName(val string) {
	s.ProcessConfiguration.Name = val
}

// SetDestination sets the value of Destination.
func (s *ConfigureProcessConfigurationRequest) SetDestination(val ConfigureProcessConfigurationRequestDestination) {
	s.ProcessConfiguration.Settings.Destination = CreateProcessConfigurationRequestDestination(val)
}

// SetParameters sets the value of Parameters.
func (s *ConfigureProcessConfigurationRequest) SetParameters(val string) {
	s.ProcessConfiguration.Settings.Parameters = val
}

type ConfigureProcessConfigurationRequestDestination string

const (
	ConfigureProcessConfigurationRequestDestinationSimplenotification ConfigureProcessConfigurationRequestDestination = "simplenotification"
	ConfigureProcessConfigurationRequestDestinationSimplemq           ConfigureProcessConfigurationRequestDestination = "simplemq"
)

// AllValues returns all ConfigureProcessConfigurationRequestDestination values.
func (ConfigureProcessConfigurationRequestDestination) AllValues() []ConfigureProcessConfigurationRequestDestination {
	return []ConfigureProcessConfigurationRequestDestination{
		ConfigureProcessConfigurationRequestDestinationSimplenotification,
		ConfigureProcessConfigurationRequestDestinationSimplemq,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ConfigureProcessConfigurationRequestDestination) MarshalText() ([]byte, error) {
	switch s {
	case ConfigureProcessConfigurationRequestDestinationSimplenotification:
		return []byte(s), nil
	case ConfigureProcessConfigurationRequestDestinationSimplemq:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ConfigureProcessConfigurationRequestDestination) UnmarshalText(data []byte) error {
	switch ConfigureProcessConfigurationRequestDestination(data) {
	case ConfigureProcessConfigurationRequestDestinationSimplenotification:
		*s = ConfigureProcessConfigurationRequestDestinationSimplenotification
		return nil
	case ConfigureProcessConfigurationRequestDestinationSimplemq:
		*s = ConfigureProcessConfigurationRequestDestinationSimplemq
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type ConfigureProcessConfigurationSecretBadRequest Error

func (*ConfigureProcessConfigurationSecretBadRequest) configureProcessConfigurationSecretRes() {}

type ConfigureProcessConfigurationSecretInternalServerError Error

func (*ConfigureProcessConfigurationSecretInternalServerError) configureProcessConfigurationSecretRes() {
}

type ConfigureProcessConfigurationSecretNotFound Error

func (*ConfigureProcessConfigurationSecretNotFound) configureProcessConfigurationSecretRes() {}

type ConfigureProcessConfigurationSecretOK struct {
	Success bool `json:is_ok`
}

// GetResult returns the value of Result.
func (s *ConfigureProcessConfigurationSecretOK) GetSuccess() bool {
	return s.Success
}

// SetResult sets the value of Result.
func (s *ConfigureProcessConfigurationSecretOK) SetSuccess(val bool) {
	s.Success = val
}

func (*ConfigureProcessConfigurationSecretOK) configureProcessConfigurationSecretRes() {}

type ProcessConfigurationSecret struct {
	AccessToken string `json:,omitempty`
	AccessTokenSecret string `json:,omitempty`
	APIKey string `json:,omitempty`
}

// Ref: #/components/schemas/ConfigureProcessConfigurationSecretRequest
type ConfigureProcessConfigurationSecretRequest struct {
	Secret ProcessConfigurationSecret `json:"Secret"`
}

// GetSecrets returns the value of Secrets.
func (s *ConfigureProcessConfigurationSecretRequest) GetSecret() ProcessConfigurationSecret {
	return s.Secret
}

// SetSecrets sets the value of Secrets.
func (s *ConfigureProcessConfigurationSecretRequest) SetSecret(val ProcessConfigurationSecret) {
	s.Secret = val
}

type ConfigureScheduleBadRequest Error

func (*ConfigureScheduleBadRequest) configureScheduleRes() {}

type ConfigureScheduleInternalServerError Error

func (*ConfigureScheduleInternalServerError) configureScheduleRes() {}

type ConfigureScheduleNotFound Error

func (*ConfigureScheduleNotFound) configureScheduleRes() {}

type ConfigureScheduleOK struct {
	Schedule Schedule `json:"CommonServiceItem"`
}

// GetResult returns the value of Result.
func (s *ConfigureScheduleOK) GetSchedule() Schedule {
	return s.Schedule
}

// SetResult sets the value of Result.
func (s *ConfigureScheduleOK) SetSchedule(val Schedule) {
	s.Schedule = val
}

func (*ConfigureScheduleOK) configureScheduleRes() {}

type ConfigureScheduleRequest struct {
	Schedule ScheduleRequestSettings `json:"CommonServiceItem"`
}

type ConfigureScheduleRequestRecurringUnit string

const (
	ConfigureScheduleRequestRecurringUnitMin  ConfigureScheduleRequestRecurringUnit = "min"
	ConfigureScheduleRequestRecurringUnitHour ConfigureScheduleRequestRecurringUnit = "hour"
	ConfigureScheduleRequestRecurringUnitDay  ConfigureScheduleRequestRecurringUnit = "day"
)

// AllValues returns all ConfigureScheduleRequestRecurringUnit values.
func (ConfigureScheduleRequestRecurringUnit) AllValues() []ConfigureScheduleRequestRecurringUnit {
	return []ConfigureScheduleRequestRecurringUnit{
		ConfigureScheduleRequestRecurringUnitMin,
		ConfigureScheduleRequestRecurringUnitHour,
		ConfigureScheduleRequestRecurringUnitDay,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ConfigureScheduleRequestRecurringUnit) MarshalText() ([]byte, error) {
	switch s {
	case ConfigureScheduleRequestRecurringUnitMin:
		return []byte(s), nil
	case ConfigureScheduleRequestRecurringUnitHour:
		return []byte(s), nil
	case ConfigureScheduleRequestRecurringUnitDay:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ConfigureScheduleRequestRecurringUnit) UnmarshalText(data []byte) error {
	switch ConfigureScheduleRequestRecurringUnit(data) {
	case ConfigureScheduleRequestRecurringUnitMin:
		*s = ConfigureScheduleRequestRecurringUnitMin
		return nil
	case ConfigureScheduleRequestRecurringUnitHour:
		*s = ConfigureScheduleRequestRecurringUnitHour
		return nil
	case ConfigureScheduleRequestRecurringUnitDay:
		*s = ConfigureScheduleRequestRecurringUnitDay
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type CreateProcessConfigurationBadRequest Error

func (*CreateProcessConfigurationBadRequest) createProcessConfigurationRes() {}

type CreateProcessConfigurationConflict Error

func (*CreateProcessConfigurationConflict) createProcessConfigurationRes() {}

type CreateProcessConfigurationInternalServerError Error

func (*CreateProcessConfigurationInternalServerError) createProcessConfigurationRes() {}

type CreateProcessConfigurationOK struct {
	ProcessConfiguration ProcessConfiguration `json:"CommonServiceItem"`
}

// GetProcess returns the value of Process.
func (s *CreateProcessConfigurationOK) GetProcessConfiguration() ProcessConfiguration {
	return s.ProcessConfiguration
}

// SetProcess sets the value of Process.
func (s *CreateProcessConfigurationOK) SetProcessConfiguration(val ProcessConfiguration) {
	s.ProcessConfiguration = val
}

func (*CreateProcessConfigurationOK) createProcessConfigurationRes() {}

type DestinationSettings struct {
	Destination CreateProcessConfigurationRequestDestination `json:"Destination"`
	Parameters  string `json:"Parameters"`
}

type ProcessConfigurationProvider struct {
	Class string `json:"Class"`
}

type ProcessConfigurationRequestSettings struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Settings    DestinationSettings`json:"Settings"`
	Provider 	ProcessConfigurationProvider `json:"Provider"`
}

// Ref: #/components/schemas/CreateProcessConfigurationRequest
type CreateProcessConfigurationRequest struct {
	ProcessConfiguration ProcessConfigurationRequestSettings `json:"CommonServiceItem"`
}

// GetName returns the value of Name.
func (s *CreateProcessConfigurationRequest) GetName() string {
	return s.ProcessConfiguration.Name
}

// GetDestination returns the value of Destination.
func (s *CreateProcessConfigurationRequest) GetSettings() DestinationSettings {
	return s.ProcessConfiguration.Settings
}

// GetParameters returns the value of Parameters.
func (s *CreateProcessConfigurationRequest) GetParameters() string {
	return s.ProcessConfiguration.Settings.Parameters
}

// SetName sets the value of Name.
func (s *CreateProcessConfigurationRequest) SetName(val string) {
	s.ProcessConfiguration.Name = val
}

// SetDestination sets the value of Destination.
func (s *CreateProcessConfigurationRequest) SetDestination(val DestinationSettings) {
	s.ProcessConfiguration.Settings = val
}

// SetParameters sets the value of Parameters.
func (s *CreateProcessConfigurationRequest) SetParameters(val string) {
	s.ProcessConfiguration.Settings.Parameters = val
}

type CreateProcessConfigurationRequestDestination string

const (
	CreateProcessConfigurationRequestDestinationSimplenotification CreateProcessConfigurationRequestDestination = "simplenotification"
	CreateProcessConfigurationRequestDestinationSimplemq           CreateProcessConfigurationRequestDestination = "simplemq"
)

// AllValues returns all CreateProcessConfigurationRequestDestination values.
func (CreateProcessConfigurationRequestDestination) AllValues() []CreateProcessConfigurationRequestDestination {
	return []CreateProcessConfigurationRequestDestination{
		CreateProcessConfigurationRequestDestinationSimplenotification,
		CreateProcessConfigurationRequestDestinationSimplemq,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s CreateProcessConfigurationRequestDestination) MarshalText() ([]byte, error) {
	switch s {
	case CreateProcessConfigurationRequestDestinationSimplenotification:
		return []byte(s), nil
	case CreateProcessConfigurationRequestDestinationSimplemq:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CreateProcessConfigurationRequestDestination) UnmarshalText(data []byte) error {
	switch CreateProcessConfigurationRequestDestination(data) {
	case CreateProcessConfigurationRequestDestinationSimplenotification:
		*s = CreateProcessConfigurationRequestDestinationSimplenotification
		return nil
	case CreateProcessConfigurationRequestDestinationSimplemq:
		*s = CreateProcessConfigurationRequestDestinationSimplemq
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type CreateProcessConfigurationUnauthorized Error

func (*CreateProcessConfigurationUnauthorized) createProcessConfigurationRes() {}

type CreateScheduleBadRequest Error

func (*CreateScheduleBadRequest) createScheduleRes() {}

type CreateScheduleConflict Error

func (*CreateScheduleConflict) createScheduleRes() {}

type CreateScheduleInternalServerError Error

func (*CreateScheduleInternalServerError) createScheduleRes() {}

type Schedule struct {
	ID    int64
	Name  string
	Description string
	Settings ScheduleSettings
	SettingsHash string
	Availability string
	CreatedAt string
	ModifiedAt string
	Status OptScheduleLastResult
	Tags []string
}

type CreateScheduleOK struct {
	Schedule Schedule `json:"CommonServiceItem"`
}

// GetSchedule returns the value of Schedule.
func (s *CreateScheduleOK) GetSchedule() Schedule {
	return s.Schedule
}

// SetSchedule sets the value of Schedule.
func (s *CreateScheduleOK) SetSchedule(val Schedule) {
	s.Schedule = val
}

func (*CreateScheduleOK) createScheduleRes() {}

type ScheduleSettings struct {
	ProcessConfigurationID string                             `json:"ProcessConfigurationID"`
	RecurringStep          int                                `json:"RecurringStep"`
	RecurringUnit          CreateScheduleRequestRecurringUnit `json:"RecurringUnit"`
	StartsAt               int64                              `json:"StartsAt,omitempty"`
}

type ScheduleProvider struct {
	Class string `json:"Class"`
}

type ScheduleRequestSettings struct {
	Name        string           `json:"Name"`
	Description string           `json:"Description"`
	Settings    ScheduleSettings `json:"Settings"`
	Provider    ScheduleProvider `json:"Provider"`
}

// Ref: #/components/schemas/CreateScheduleRequest
type CreateScheduleRequest struct {
	Schedule ScheduleRequestSettings `json:"CommonServiceItem"`
}

// GetName returns the value of Name.
func (s *CreateScheduleRequest) GetName() string {
	return s.Schedule.Name
}

// GetProcessConfigurationID returns the value of ProcessConfigurationID.
func (s *CreateScheduleRequest) GetProcessConfigurationID() string {
	return s.Schedule.Settings.ProcessConfigurationID
}

// GetRecurringStep returns the value of RecurringStep.
func (s *CreateScheduleRequest) GetRecurringStep() int {
	return s.Schedule.Settings.RecurringStep
}

// GetRecurringUnit returns the value of RecurringUnit.
func (s *CreateScheduleRequest) GetRecurringUnit() CreateScheduleRequestRecurringUnit {
	return s.Schedule.Settings.RecurringUnit
}

// GetStartsAt returns the value of StartsAt.
func (s *CreateScheduleRequest) GetStartsAt() int64 {
	return s.Schedule.Settings.StartsAt
}

// SetName sets the value of Name.
func (s *CreateScheduleRequest) SetName(val string) {
	s.Schedule.Name = val
}

// SetProcessConfigurationID sets the value of ProcessConfigurationID.
func (s *CreateScheduleRequest) SetProcessConfigurationID(val string) {
	s.Schedule.Settings.ProcessConfigurationID = val
}

// SetRecurringStep sets the value of RecurringStep.
func (s *CreateScheduleRequest) SetRecurringStep(val int) {
	s.Schedule.Settings.RecurringStep = val
}

// SetRecurringUnit sets the value of RecurringUnit.
func (s *CreateScheduleRequest) SetRecurringUnit(val CreateScheduleRequestRecurringUnit) {
	s.Schedule.Settings.RecurringUnit = val
}

// SetStartsAt sets the value of StartsAt.
func (s *CreateScheduleRequest) SetStartsAt(val int64) {
	s.Schedule.Settings.StartsAt = val
}

type CreateScheduleRequestRecurringUnit string

const (
	CreateScheduleRequestRecurringUnitMin  CreateScheduleRequestRecurringUnit = "min"
	CreateScheduleRequestRecurringUnitHour CreateScheduleRequestRecurringUnit = "hour"
	CreateScheduleRequestRecurringUnitDay  CreateScheduleRequestRecurringUnit = "day"
)

// AllValues returns all CreateScheduleRequestRecurringUnit values.
func (CreateScheduleRequestRecurringUnit) AllValues() []CreateScheduleRequestRecurringUnit {
	return []CreateScheduleRequestRecurringUnit{
		CreateScheduleRequestRecurringUnitMin,
		CreateScheduleRequestRecurringUnitHour,
		CreateScheduleRequestRecurringUnitDay,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s CreateScheduleRequestRecurringUnit) MarshalText() ([]byte, error) {
	switch s {
	case CreateScheduleRequestRecurringUnitMin:
		return []byte(s), nil
	case CreateScheduleRequestRecurringUnitHour:
		return []byte(s), nil
	case CreateScheduleRequestRecurringUnitDay:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CreateScheduleRequestRecurringUnit) UnmarshalText(data []byte) error {
	switch CreateScheduleRequestRecurringUnit(data) {
	case CreateScheduleRequestRecurringUnitMin:
		*s = CreateScheduleRequestRecurringUnitMin
		return nil
	case CreateScheduleRequestRecurringUnitHour:
		*s = CreateScheduleRequestRecurringUnitHour
		return nil
	case CreateScheduleRequestRecurringUnitDay:
		*s = CreateScheduleRequestRecurringUnitDay
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type CreateScheduleUnauthorized Error

func (*CreateScheduleUnauthorized) createScheduleRes() {}

type DeleteProcessConfigurationBadRequest Error

func (*DeleteProcessConfigurationBadRequest) deleteProcessConfigurationRes() {}

type DeleteProcessConfigurationConflict Error

func (*DeleteProcessConfigurationConflict) deleteProcessConfigurationRes() {}

type DeleteProcessConfigurationInternalServerError Error

func (*DeleteProcessConfigurationInternalServerError) deleteProcessConfigurationRes() {}

type DeleteProcessConfigurationNotFound Error

func (*DeleteProcessConfigurationNotFound) deleteProcessConfigurationRes() {}

type DeleteProcessConfigurationOK struct {
	ProcessConfiguration ProcessConfiguration `json:"CommonServiceItem"`
}

// GetResult returns the value of Result.
func (s *DeleteProcessConfigurationOK) GetProcessConfiguration() ProcessConfiguration {
	return s.ProcessConfiguration
}

// SetResult sets the value of Result.
func (s *DeleteProcessConfigurationOK) SetProcessConfiguration(val ProcessConfiguration) {
	s.ProcessConfiguration = val
}

func (*DeleteProcessConfigurationOK) deleteProcessConfigurationRes() {}

type DeleteProcessConfigurationUnauthorized Error

func (*DeleteProcessConfigurationUnauthorized) deleteProcessConfigurationRes() {}

type DeleteScheduleBadRequest Error

func (*DeleteScheduleBadRequest) deleteScheduleRes() {}

type DeleteScheduleInternalServerError Error

func (*DeleteScheduleInternalServerError) deleteScheduleRes() {}

type DeleteScheduleNotFound Error

func (*DeleteScheduleNotFound) deleteScheduleRes() {}

type DeleteScheduleOK struct {
	Schedule Schedule `json:"CommonServiceItem"`
}

// GetResult returns the value of Result.
func (s *DeleteScheduleOK) GetSchedule() Schedule {
	return s.Schedule
}

// SetResult sets the value of Result.
func (s *DeleteScheduleOK) SetSchedule(val Schedule) {
	s.Schedule = val
}

func (*DeleteScheduleOK) deleteScheduleRes() {}

type DeleteScheduleUnauthorized Error

func (*DeleteScheduleUnauthorized) deleteScheduleRes() {}

// Error content.
// Ref: #/components/schemas/Error
type Error struct {
	Code    OptString `json:"error_code"`
	Message OptString `json:"error_msg`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() OptString {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() OptString {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val OptString) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val OptString) {
	s.Message = val
}

type GetProcessConfigurationByIdBadRequest Error

func (*GetProcessConfigurationByIdBadRequest) getProcessConfigurationByIdRes() {}

type GetProcessConfigurationByIdInternalServerError Error

func (*GetProcessConfigurationByIdInternalServerError) getProcessConfigurationByIdRes() {}

type GetProcessConfigurationByIdNotFound Error

func (*GetProcessConfigurationByIdNotFound) getProcessConfigurationByIdRes() {}

type ProcessConfiguration struct {
	Index int64
	ID    int64
	Name  string
	Description string
	Settings DestinationSettings
	SettingsHash string
	Availability string
	CreatedAt string
	ModifiedAt string
	Tags []string
}

type GetProcessConfigurationByIdOK struct {
	ProcessConfiguration ProcessConfiguration `json:"CommonServiceItem"`
}

// GetProcessConfiguration returns the value of ProcessConfiguration.
func (s *GetProcessConfigurationByIdOK) GetProcessConfiguration() ProcessConfiguration {
	return s.ProcessConfiguration
}

// SetProcessConfiguration sets the value of ProcessConfiguration.
func (s *GetProcessConfigurationByIdOK) SetProcessConfiguration(val ProcessConfiguration) {
	s.ProcessConfiguration = val
}

func (*GetProcessConfigurationByIdOK) getProcessConfigurationByIdRes() {}

type GetProcessConfigurationByIdUnauthorized Error

func (*GetProcessConfigurationByIdUnauthorized) getProcessConfigurationByIdRes() {}

type GetProcessConfigurationsBadRequest Error

func (*GetProcessConfigurationsBadRequest) getProcessConfigurationsByAccountRes() {}

type GetProcessConfigurationsInternalServerError Error

func (*GetProcessConfigurationsInternalServerError) getProcessConfigurationsByAccountRes() {}

type GetProcessConfigurationsOK struct {
	From  int64
	Count int64
	Total int64
	ProcessConfigurations []ProcessConfiguration `json:"CommonServiceItems"`
}

// GetCount returns the value of Count.
func (s *GetProcessConfigurationsOK) GetCount() int64 {
	return s.Count
}

// GetTotal returns the value of Total.
func (s *GetProcessConfigurationsOK) GetTotal() int64 {
	return s.Total
}

// GetProcessConfigurations returns the value of ProcessConfigurations.
func (s *GetProcessConfigurationsOK) GetProcessConfigurations() []ProcessConfiguration {
	return s.ProcessConfigurations
}

// SetCount sets the value of Count.
func (s *GetProcessConfigurationsOK) SetCount(val int64) {
	s.Count = val
}

// SetTotal sets the value of Total.
func (s *GetProcessConfigurationsOK) SetTotal(val int64) {
	s.Total = val
}

// SetProcessConfigurations sets the value of ProcessConfigurations.
func (s *GetProcessConfigurationsOK) SetProcessConfigurations(val []ProcessConfiguration) {
	s.ProcessConfigurations = val
}

func (*GetProcessConfigurationsOK) getProcessConfigurationsByAccountRes() {}

type GetProcessConfigurationsUnauthorized Error

func (*GetProcessConfigurationsUnauthorized) getProcessConfigurationsByAccountRes() {}

type GetScheduleByIdBadRequest Error

func (*GetScheduleByIdBadRequest) getScheduleByIdRes() {}

type GetScheduleByIdInternalServerError Error

func (*GetScheduleByIdInternalServerError) getScheduleByIdRes() {}

type GetScheduleByIdNotFound Error

func (*GetScheduleByIdNotFound) getScheduleByIdRes() {}

type GetScheduleByIdOK struct {
	Schedule Schedule `json:"CommonServiceItem"`
}

// GetSchedule returns the value of Schedule.
func (s *GetScheduleByIdOK) GetSchedule() Schedule {
	return s.Schedule
}

// SetSchedule sets the value of Schedule.
func (s *GetScheduleByIdOK) SetSchedule(val Schedule) {
	s.Schedule = val
}

func (*GetScheduleByIdOK) getScheduleByIdRes() {}

type GetScheduleByIdUnauthorized Error

func (*GetScheduleByIdUnauthorized) getScheduleByIdRes() {}

type GetSchedulesBadRequest Error

func (*GetSchedulesBadRequest) getSchedulesByAccountRes() {}

type GetSchedulesInternalServerError Error

func (*GetSchedulesInternalServerError) getSchedulesByAccountRes() {}

type GetSchedulesOK struct {
	From  int64
	Count int64
	Total int64
	Schedules []Schedule `json:"CommonServiceItems"`
}

// GetSchedules returns the value of Schedules.
func (s *GetSchedulesOK) GetSchedules() []Schedule {
	return s.Schedules
}

// SetSchedules sets the value of Schedules.
func (s *GetSchedulesOK) SetSchedules(val []Schedule) {
	s.Schedules = val
}

func (*GetSchedulesOK) getSchedulesByAccountRes() {}

type GetSchedulesUnauthorized Error

func (*GetSchedulesUnauthorized) getSchedulesByAccountRes() {}

// NewNilString returns new NilString with value set to v.
func NewNilString(v string) NilString {
	return NilString{
		Value: v,
	}
}

// NilString is nullable string.
type NilString struct {
	Value string
	Null  bool
}

// SetTo sets value to v.
func (o *NilString) SetTo(v string) {
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o NilString) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *NilString) SetToNull() {
	o.Null = true
	var v string
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o NilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptConfigureScheduleRequestRecurringUnit returns new OptConfigureScheduleRequestRecurringUnit with value set to v.
func NewOptConfigureScheduleRequestRecurringUnit(v ConfigureScheduleRequestRecurringUnit) OptConfigureScheduleRequestRecurringUnit {
	return OptConfigureScheduleRequestRecurringUnit{
		Value: v,
		Set:   true,
	}
}

// OptConfigureScheduleRequestRecurringUnit is optional ConfigureScheduleRequestRecurringUnit.
type OptConfigureScheduleRequestRecurringUnit struct {
	Value ConfigureScheduleRequestRecurringUnit
	Set   bool
}

// IsSet returns true if OptConfigureScheduleRequestRecurringUnit was set.
func (o OptConfigureScheduleRequestRecurringUnit) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptConfigureScheduleRequestRecurringUnit) Reset() {
	var v ConfigureScheduleRequestRecurringUnit
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptConfigureScheduleRequestRecurringUnit) SetTo(v ConfigureScheduleRequestRecurringUnit) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptConfigureScheduleRequestRecurringUnit) Get() (v ConfigureScheduleRequestRecurringUnit, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptConfigureScheduleRequestRecurringUnit) Or(d ConfigureScheduleRequestRecurringUnit) ConfigureScheduleRequestRecurringUnit {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCreateScheduleRequestRecurringUnit returns new OptCreateScheduleRequestRecurringUnit with value set to v.
func NewOptCreateScheduleRequestRecurringUnit(v CreateScheduleRequestRecurringUnit) OptCreateScheduleRequestRecurringUnit {
	return OptCreateScheduleRequestRecurringUnit{
		Value: v,
		Set:   true,
	}
}

// OptCreateScheduleRequestRecurringUnit is optional CreateScheduleRequestRecurringUnit.
type OptCreateScheduleRequestRecurringUnit struct {
	Value CreateScheduleRequestRecurringUnit
	Set   bool
}

// IsSet returns true if OptCreateScheduleRequestRecurringUnit was set.
func (o OptCreateScheduleRequestRecurringUnit) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCreateScheduleRequestRecurringUnit) Reset() {
	var v CreateScheduleRequestRecurringUnit
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCreateScheduleRequestRecurringUnit) SetTo(v CreateScheduleRequestRecurringUnit) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCreateScheduleRequestRecurringUnit) Get() (v CreateScheduleRequestRecurringUnit, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCreateScheduleRequestRecurringUnit) Or(d CreateScheduleRequestRecurringUnit) CreateScheduleRequestRecurringUnit {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt64 returns new OptInt64 with value set to v.
func NewOptInt64(v int64) OptInt64 {
	return OptInt64{
		Value: v,
		Set:   true,
	}
}

// OptInt64 is optional int64.
type OptInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if OptInt64 was set.
func (o OptInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptProcessConfiguration returns new OptProcessConfiguration with value set to v.
func NewOptProcessConfiguration(v ProcessConfiguration) OptProcessConfiguration {
	return OptProcessConfiguration{
		Value: v,
		Set:   true,
	}
}

// OptProcessConfiguration is optional ProcessConfiguration.
type OptProcessConfiguration struct {
	Value ProcessConfiguration
	Set   bool
}

// IsSet returns true if OptProcessConfiguration was set.
func (o OptProcessConfiguration) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptProcessConfiguration) Reset() {
	var v ProcessConfiguration
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptProcessConfiguration) SetTo(v ProcessConfiguration) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptProcessConfiguration) Get() (v ProcessConfiguration, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptProcessConfiguration) Or(d ProcessConfiguration) ProcessConfiguration {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSchedule returns new OptSchedule with value set to v.
func NewOptSchedule(v Schedule) OptSchedule {
	return OptSchedule{
		Value: v,
		Set:   true,
	}
}

// OptSchedule is optional Schedule.
type OptSchedule struct {
	Value Schedule
	Set   bool
}

// IsSet returns true if OptSchedule was set.
func (o OptSchedule) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSchedule) Reset() {
	var v Schedule
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSchedule) SetTo(v Schedule) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSchedule) Get() (v Schedule, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSchedule) Or(d Schedule) Schedule {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptScheduleLastResult returns new OptScheduleLastResult with value set to v.
func NewOptScheduleLastResult(v ScheduleLastResult) OptScheduleLastResult {
	return OptScheduleLastResult{
		Value: v,
		Set:   true,
	}
}

// OptScheduleLastResult is optional ScheduleLastResult.
type OptScheduleLastResult struct {
	Value ScheduleLastResult
	Set   bool
}

// IsSet returns true if OptScheduleLastResult was set.
func (o OptScheduleLastResult) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptScheduleLastResult) Reset() {
	var v ScheduleLastResult
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptScheduleLastResult) SetTo(v ScheduleLastResult) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptScheduleLastResult) Get() (v ScheduleLastResult, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptScheduleLastResult) Or(d ScheduleLastResult) ScheduleLastResult {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}


type ProcessConfigurationDestination string

const (
	ProcessConfigurationDestinationSimplenotification ProcessConfigurationDestination = "simplenotification"
	ProcessConfigurationDestinationSimplemq           ProcessConfigurationDestination = "simplemq"
)

// AllValues returns all ProcessConfigurationDestination values.
func (ProcessConfigurationDestination) AllValues() []ProcessConfigurationDestination {
	return []ProcessConfigurationDestination{
		ProcessConfigurationDestinationSimplenotification,
		ProcessConfigurationDestinationSimplemq,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ProcessConfigurationDestination) MarshalText() ([]byte, error) {
	switch s {
	case ProcessConfigurationDestinationSimplenotification:
		return []byte(s), nil
	case ProcessConfigurationDestinationSimplemq:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ProcessConfigurationDestination) UnmarshalText(data []byte) error {
	switch ProcessConfigurationDestination(data) {
	case ProcessConfigurationDestinationSimplenotification:
		*s = ProcessConfigurationDestinationSimplenotification
		return nil
	case ProcessConfigurationDestinationSimplemq:
		*s = ProcessConfigurationDestinationSimplemq
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/ScheduleLastResult
type ScheduleLastResult struct {
	Success   bool   `json:"Success"`
	Message   string `json:"Message"`
	UpdatedAt int64  `json:"UpdatedAt"`
}

// GetSuccess returns the value of Success.
func (s *ScheduleLastResult) GetSuccess() bool {
	return s.Success
}

// GetMessage returns the value of Message.
func (s *ScheduleLastResult) GetMessage() string {
	return s.Message
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *ScheduleLastResult) GetUpdatedAt() int64 {
	return s.UpdatedAt
}

// SetSuccess sets the value of Success.
func (s *ScheduleLastResult) SetSuccess(val bool) {
	s.Success = val
}

// SetMessage sets the value of Message.
func (s *ScheduleLastResult) SetMessage(val string) {
	s.Message = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *ScheduleLastResult) SetUpdatedAt(val int64) {
	s.UpdatedAt = val
}

type ScheduleRecurringUnit string

const (
	ScheduleRecurringUnitMin  ScheduleRecurringUnit = "min"
	ScheduleRecurringUnitHour ScheduleRecurringUnit = "hour"
	ScheduleRecurringUnitDay  ScheduleRecurringUnit = "day"
)

// AllValues returns all ScheduleRecurringUnit values.
func (ScheduleRecurringUnit) AllValues() []ScheduleRecurringUnit {
	return []ScheduleRecurringUnit{
		ScheduleRecurringUnitMin,
		ScheduleRecurringUnitHour,
		ScheduleRecurringUnitDay,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ScheduleRecurringUnit) MarshalText() ([]byte, error) {
	switch s {
	case ScheduleRecurringUnitMin:
		return []byte(s), nil
	case ScheduleRecurringUnitHour:
		return []byte(s), nil
	case ScheduleRecurringUnitDay:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ScheduleRecurringUnit) UnmarshalText(data []byte) error {
	switch ScheduleRecurringUnit(data) {
	case ScheduleRecurringUnitMin:
		*s = ScheduleRecurringUnitMin
		return nil
	case ScheduleRecurringUnitHour:
		*s = ScheduleRecurringUnitHour
		return nil
	case ScheduleRecurringUnitDay:
		*s = ScheduleRecurringUnitDay
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}
