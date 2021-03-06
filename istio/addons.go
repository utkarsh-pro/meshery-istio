package istio

import (
	"context"
	"fmt"
	"strings"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/status"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// installAddon installs/uninstalls an addon in the given namespace
//
// the template defines the manifest's link/location which needs to be used to
// install the addon
func (istio *Istio) installAddon(namespace string, del bool, service string, patch string, templates []adapter.Template) (string, error) {
	st := status.Installing

	if del {
		st = status.Removing
	}

	istio.Log.Debug(fmt.Sprintf("Overidden namespace: %s", namespace))
	namespace = ""

	for _, template := range templates {
		contents, err := readFileSource(string(template))
		if err != nil {
			return st, ErrAddonFromTemplate(err)
		}

		err = istio.applyManifest([]byte(contents), del, namespace)
		// Specifically choosing to ignore kiali dashboard's error.
		// Referring to: https://github.com/kiali/kiali/issues/3112
		if err != nil && !strings.Contains(err.Error(), "no matches for kind \"MonitoringDashboard\" in version \"monitoring.kiali.io/v1alpha1\"") {
			return st, ErrAddonFromTemplate(err)
		}
	}

	jsonContents, err := readFileSource(patch)
	if err != nil {
		return st, ErrAddonFromTemplate(err)
	}

	_, err = istio.KubeClient.CoreV1().Services("istio-system").Patch(context.TODO(), service, types.MergePatchType, []byte(jsonContents), metav1.PatchOptions{})
	if err != nil {
		return st, ErrAddonFromTemplate(err)
	}

	return status.Installed, nil
}
