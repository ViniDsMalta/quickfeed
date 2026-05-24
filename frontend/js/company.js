const params =
  new URLSearchParams(window.location.search)

const slug = params.get("slug")

async function loadCompany() {

  if (!slug) {

    show("missing slug")

    return
  }

  const response = await fetch(
    API + "/companies/" + slug
  )

  if (!response.ok) {

    show("company not found")

    return
  }

  const company = await response.json()

  document.getElementById("companyName")
    .textContent = company.name

  document.getElementById("companySlug")
    .textContent = "@" + company.slug
}

async function sendFeedback() {

  const message =
    document.getElementById("feedbackMessage").value

  const response = await fetch(
    API + "/companies/" + slug + "/feedback",
    {
      method: "POST",

      headers: {
        "Content-Type": "application/json"
      },

      body: JSON.stringify({
        message
      })
    }
  )

  const data = await response.text()

  show(data)
}

function show(text) {

  document.getElementById("output")
    .textContent = text
}

loadCompany()