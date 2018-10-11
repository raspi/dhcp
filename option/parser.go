package option

import (
	"./base"
	"./defaults"
)

func RawToTypes(rawOptions []OptionRaw, typeMap defaults.TypeMap) (opts []base.OptionBase, err error) {

	for _, o := range rawOptions {
		funcRef, ok := typeMap[o.Type]

		if !ok {
			var unknown OptUnknown
			base, err := unknown.Read(o.Value)

			if err != nil {
				return opts, err
			}
			base.Type = o.Type

			opts = append(opts, base)
			continue
		}

		a, err := funcRef(o.Value)
		if err != nil {
			return opts, err
		}

		a.Type = o.Type

		opts = append(opts, a)
	}

	return opts, nil
}
