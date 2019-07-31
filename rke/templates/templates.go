package templates

/*
Not including Vsphere(cloudProvider) and Authz templates
Will they change and require Rancher to pass them to RKE
depending on Kubernetes version?
*/

const (
	Calico        = "calico"
	Canal         = "canal"
	Flannel       = "flannel"
	Weave         = "weave"
	CoreDNS       = "coreDNS"
	KubeDNS       = "kubeDNS"
	MetricsServer = "metricsServer"
	NginxIngress  = "nginxIngress"
	TemplateKeys  = "templateKeys"

	calicov112 = "calico-v1.12"
	calicov113 = "calico-v1.13"
	calicov115 = "calico-v1.15"

	canalv115 = "canal-v1.15"
	canalv113 = "canal-v1.13"
	canalv112 = "canal-v1.12"

	flannelv115    = "flannel-v1.15"
	flannelDefault = "flannel-default"

	coreDnsDefault = "coredns-default"
	kubeDnsDefault = "kubedns-default"

	metricsServerDefault = "metricsserver-default"

	weaveDefault        = "weave-default"
	nginxIngressDefault = "nginxingress-default"
)

func LoadK8sVersionedTemplates() map[string]map[string]string {
	return map[string]map[string]string{
		Calico: {
			">=1.15.0":         calicov115,
			">=1.13.0 <1.15.0": calicov113,
			"default":          calicov112,
		},
		Canal: {
			">=1.15.0":         canalv115,
			">=1.13.0 <1.15.0": canalv113,
			"default":          canalv112,
		},
		Flannel: {
			">=1.15.0": flannelv115,
			"default":  flannelDefault,
		},
		CoreDNS: {
			"default": coreDnsDefault,
		},
		KubeDNS: {
			"default": kubeDnsDefault,
		},
		MetricsServer: {
			"default": metricsServerDefault,
		},
		Weave: {
			"default": weaveDefault,
		},
		NginxIngress: {
			"default": nginxIngressDefault,
		},
		TemplateKeys: getTemplates(),
	}
}

func getTemplates() map[string]string {
	return map[string]string{
		calicov113: CalicoTemplateV113,
		calicov115: CalicoTemplateV115,
		calicov112: CalicoTemplateV112,

		flannelv115:    FlannelTemplateV115,
		flannelDefault: FlannelTemplate,

		canalv113: CanalTemplateV113,
		canalv112: CanalTemplateV112,
		canalv115: CanalTemplateV115,

		coreDnsDefault: CoreDNSTemplate,
		kubeDnsDefault: KubeDNSTemplate,

		metricsServerDefault: MetricsServerTemplate,

		weaveDefault: WeaveTemplate,

		nginxIngressDefault: NginxIngressTemplate,
	}
}
