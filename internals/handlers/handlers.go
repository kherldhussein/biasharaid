package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kh3rld/biasharaid/blockchain"
	"github.com/kh3rld/biasharaid/internals/renders"
)

var data renders.FormData

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := renders.FormData{
		Title: "Welcome to BiasharaID - Your Secure Blockchain Identity Verification",
	}
	renders.RenderTemplate(w, "home.page.html", &data)
}

func Verification(w http.ResponseWriter, r *http.Request) {
	data := renders.FormData{
		Title: "Verification - BiasharaID",
	}
	renders.RenderTemplate(w, "verify.page.html", &data)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	data := renders.FormData{
		Title: "Contact Us - BiasharaID",
	}
	renders.RenderTemplate(w, "contact.page.html", &data)
}

func Details(w http.ResponseWriter, r *http.Request) {
	// Add your implementation and title here
	data := renders.FormData{
		Title: "Details - BiasharaID",
	}
	renders.RenderTemplate(w, "details.page.html", &data)
}

func DummyHandler(w http.ResponseWriter, r *http.Request) {
	// resp := blockchain.BlockchainInstance.Blocks
	data := renders.FormData{
		Title: "Dummy Data - BiasharaID",
	}
	renders.RenderTemplate(w, "dummy.page.html", &data)
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := renders.FormData{
			Title: "Verify Your Identity - BiasharaID",
		}
		renders.RenderTemplate(w, "verify.page.html", &data)
		return
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}

		nationalID := r.FormValue("national_id")

		if nationalID == "" {
			BadRequestHandler(w, r)
			return
		}

		var block *blockchain.Block
		for _, b := range blockchain.BlockchainInstance.Blocks {
			if b.Data.NationalID == nationalID {
				block = b
				break
			}
		}

		if block == nil {
			data := renders.FormData{
				Title: "Not Found - BiasharaID",
			}
			renders.RenderTemplate(w, "not_found.page.html", &data)
			return
		}

		// data := renders.FormData{
		// 	Title: "Verification Result - BiasharaID",
		// }
		renders.RenderTemplate(w, "verify.page.html", block)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	data := renders.FormData{
		Title: "Add Entrepreneur - BiasharaID",
	}
	renders.RenderTemplate(w, "add.page.html", &data)
}

func Addpage(w http.ResponseWriter, r *http.Request) {
	var entrepreneur blockchain.Entrepreneur
	if err := json.NewDecoder(r.Body).Decode(&entrepreneur); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if entrepreneur.FirstName == "" || entrepreneur.SecondName == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	blockchain.BlockchainInstance.AddBlock(entrepreneur)
	w.WriteHeader(http.StatusOK)
	data := renders.FormData{
		Body:  "Data Added Successfully",
		Title: "Add Entrepreneur - BiasharaID",
	}
	renders.RenderTemplate(w, "add.page.html", &data)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	data := renders.FormData{
		Title: "Page Not Found - BiasharaID",
	}
	renders.RenderTemplate(w, "404.page.html", &data)
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	data := renders.FormData{
		Title: "Bad Request - BiasharaID",
	}
	renders.RenderTemplate(w, "400.page.html", &data)
}

func ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	data := renders.FormData{
		Title: "Server Error - BiasharaID",
	}
	renders.RenderTemplate(w, "500.page.html", &data)
}
