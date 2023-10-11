package toot

type APIFunc[In any, Out any] func(args In) (Out, error)
