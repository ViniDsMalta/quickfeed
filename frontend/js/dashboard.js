const token = localStorage.getItem("token")


if (!token) {
  window.location.href = "login.html"
}

async function createCompany() {

  const name =
    document.getElementById("companyName").value

  const response = await fetch(
    API + "/companies",
    {
      method: "POST",

      headers: {
        "Content-Type": "application/json",

        "Authorization":
          "Bearer " + token
      },

      body: JSON.stringify({
        name
      })
    }
  )

  const data = await response.text()

  show(data)

  loadCompanies()
}

async function loadCompanies() {

  const response = await fetch(
    API + "/companies",
    {
      headers: {
        "Authorization":
          "Bearer " + token
      }
    }
  )

  const companies = await response.json()

  const container =
    document.getElementById("companies")

  container.innerHTML = ""

  companies.forEach(company => {

    const div = document.createElement("div")

    div.innerHTML = `
      <h3>${company.name}</h3>

      <p>
        Slug:
        ${company.slug}
      </p>

      <a
        href="company.html?slug=${company.slug}"
      >
        Public Page
      </a>

      <br><br>

      <button onclick="loadFeedbacks('${company.slug}')">
        View Feedbacks
      </button>

      <hr>
    `

    container.appendChild(div)
  })
}

function logout() {

  localStorage.removeItem("token")

  window.location.href = "login.html"
}

function show(text) {

  document.getElementById("output")
    .textContent = text
}
async function loadFeedbacks(slug) {

  const response = await fetch(
    API + "/companies/" + slug + "/feedbacks",
    {
      headers: {
        "Authorization":
          "Bearer " + token
      }
    }
  )

  const feedbacks = await response.json()

  const container =
    document.getElementById("feedbacks")

  container.innerHTML = ""

  if (feedbacks.length === 0) {

    container.innerHTML =
      "<p>No feedbacks yet</p>"

    return
  }

  feedbacks.forEach(feedback => {

    const div = document.createElement("div")

    div.innerHTML = `
      <p>${feedback.message}</p>

      <small>
        ${feedback.created_at}
      </small>

      <hr>
    `

    container.appendChild(div)
  })
}
loadCompanies()

