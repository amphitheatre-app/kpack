package volume

import (
	"context"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

type Resolver struct {
}

func (*Resolver) Resolve(ctx context.Context, sourceResolver *buildapi.SourceResolver) (corev1alpha1.ResolvedSourceConfig, error) {
	return corev1alpha1.ResolvedSourceConfig{
		Volume: &corev1alpha1.ResolvedVolumeSource{
			ClaimName: sourceResolver.Spec.Source.Volume.ClaimName,
		},
	}, nil
}

func (*Resolver) CanResolve(sourceResolver *buildapi.SourceResolver) bool {
	return sourceResolver.IsVolume()
}
