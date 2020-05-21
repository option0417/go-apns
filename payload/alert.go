package payload

type Alert struct {
	title           string
	subtitle        string
	body            string
	launchImage     string
	titleLocKey     string
	titleLocArgs    []string
	subTitleLocKey  string
	subTitleLocArgs []string
	locKey          string
	locArgs         []string
}
