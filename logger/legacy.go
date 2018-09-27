package logger

type LegacySpec struct {
	DebugOut string `default:"STDOUT" envconfig:"DEBUG_OUT"`
	InfoOut  string `default:"STDOUT" envconfig:"INFO_OUT"`
	Level    int8   `default:"2"`

	FlagsDefault string `envconfig:"FLAGS_DEFAULT"`
	FlagsInfo    string `envconfig:"FLAGS_INFO"`
	FlagsDebug   string `envconfig:"FLAGS_DEBUG"`
}
