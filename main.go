func redirectToawdev(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "https://wahyu9kdl.github.io/profile", http.StatusSeeOther)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", redirectToFreshman)

	http.ListenAndServe(":6000", mux)
}
