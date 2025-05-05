package cmd

import (
	"fmt"

	"github.com/nobbmaestro/beam/pkg/sender"
	"github.com/spf13/cobra"
)

type Options struct {
	addressDst     string
	addressSrc     string
	channels       uint16
	duration       uint64
	fps            uint32
	frequency      float64
	intensityLower uint8
	intensityUpper uint8
	priority       uint8
	universes      []uint
}

var (
	opts        = Options{}
	defaultOpts = sender.Default()
)

var rootCmd = &cobra.Command{
	Short:   "beam",
	Long:    "Command-line utility for generating and broadcasting DMX over sACN",
	PreRunE: validateOptions,
	RunE:    runRoot,
}

func validateOptions(cmd *cobra.Command, args []string) error {
	if opts.intensityUpper > 255 {
		return fmt.Errorf("")
	}
	if opts.intensityLower > opts.intensityUpper {
		return fmt.Errorf("")
	}
	if opts.channels > 512 {
		return fmt.Errorf("")
	}
	return nil
}

func runRoot(cmd *cobra.Command, args []string) error {
	s := sender.Sender{
		AddressDst:     opts.addressDst,
		AddressSrc:     opts.addressSrc,
		Channels:       opts.channels,
		Duration:       opts.duration,
		Fps:            opts.fps,
		Frequency:      opts.frequency,
		IntesintyLower: opts.intensityLower,
		IntesintyUpper: opts.intensityUpper,
		Priority:       opts.priority,
		Universes:      opts.universes,
	}
	return s.Run()
}

func SetVersion(version, commit, date string) {
	rootCmd.Version = version
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	flags := rootCmd.Flags()

	flags.StringVar(
		&opts.addressSrc,
		"src",
		defaultOpts.AddressSrc,
		"IP address of the DMX source",
	)
	flags.StringVar(
		&opts.addressDst,
		"dst",
		defaultOpts.AddressSrc,
		"IP address of the DMX destination",
	)
	flags.Uint16VarP(
		&opts.channels,
		"channels",
		"c",
		defaultOpts.Channels,
		"DMX channels at universe to send to",
	)
	flags.Uint32Var(
		&opts.fps,
		"fps",
		defaultOpts.Fps,
		"frames per second per universe",
	)
	flags.Float64VarP(
		&opts.frequency,
		"frequency",
		"f",
		defaultOpts.Frequency,
		"frequency of the generated signal",
	)
	flags.Uint8VarP(
		&opts.intensityUpper,
		"upper",
		"i",
		defaultOpts.IntesintyUpper,
		"DMX channels upper output intensity",
	)
	flags.Uint8VarP(
		&opts.intensityLower,
		"lower",
		"I",
		defaultOpts.IntesintyLower,
		"DMX channels lower output intensity",
	)
	flags.UintSliceVarP(
		&opts.universes,
		"universes",
		"u",
		defaultOpts.Universes,
		"sACN universe(s) to send to",
	)
	flags.Uint64VarP(
		&opts.duration,
		"duration",
		"d",
		defaultOpts.Duration,
		"broadcast duration in seconds",
	)
	flags.Uint8VarP(
		&opts.priority,
		"priority",
		"p",
		defaultOpts.Priority,
		"sACN stream priority",
	)

	rootCmd.MarkFlagRequired("src")
}
