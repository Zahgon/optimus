package specio

type jobSpecReadWriterOpt func(*jobSpecReadWriter) error

func WithJobSpecParentReading() jobSpecReadWriterOpt {
	_ = "STUB: not implemented"
	return *new(jobSpecReadWriterOpt)
}
