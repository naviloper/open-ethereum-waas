package di

import (
	"github.com/sarulabs/di/v2"
)

var builder *di.Builder

func NewContainer() di.Container {
	if builder == nil {
		//todo handle builder error
		builder, _ = di.NewBuilder( di.App )
	}

	//Add services to builder

	//return container from builder
	return builder.Build()
}

