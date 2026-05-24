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

  if (companies.length === 0) {
    container.innerHTML = `<p style="color: var(--text-muted); font-style: italic; text-align: center; padding: 20px 0;">No companies created yet.</p>`
    return
  }

  companies.forEach(company => {

    const div = document.createElement("div")
    div.className = "company-item"

    div.innerHTML = `
      <div class="company-item-header">
        <h3>${company.name}</h3>
        <span class="badge badge-slug">@${company.slug}</span>
      </div>

      <div class="company-item-actions">
        <a href="company.html?slug=${company.slug}" target="_blank">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"></path>
            <polyline points="15 3 21 3 21 9"></polyline>
            <line x1="10" y1="14" x2="21" y2="3"></line>
          </svg>
          Public Page
        </a>

        <button onclick="loadFeedbacks('${company.slug}')" class="secondary">
          View Feedbacks
        </button>
      </div>
    `

    container.appendChild(div)
  })
}

function logout() {

  localStorage.removeItem("token")

  window.location.href = "login.html"
}

function show(text) {
  const output = document.getElementById("output");
  output.textContent = text;
  output.className = ""; // clear previous alert styles
  
  const textLower = text.toLowerCase();
  if (
    textLower.includes("fail") ||
    textLower.includes("error") ||
    textLower.includes("missing") ||
    textLower.includes("not found") ||
    textLower.includes("invalid") ||
    textLower.includes("exists") ||
    textLower.includes("incorrect")
  ) {
    output.classList.add("alert-error");
  } else {
    output.classList.add("alert-success");
  }
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
      `<p style="color: var(--text-muted); font-style: italic; text-align: center; padding: 20px 0;">No feedbacks yet</p>`

    return
  }

  feedbacks.forEach(feedback => {

    const div = document.createElement("div")
    div.className = "feedback-bubble"

    let dateStr = feedback.created_at;
    try {
      const d = new Date(feedback.created_at);
      if (!isNaN(d.getTime())) {
        dateStr = d.toLocaleDateString() + ' ' + d.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'});
      }
    } catch(e) {}

    div.innerHTML = `
      <p>${feedback.message}</p>
      <small>${dateStr}</small>
    `

    container.appendChild(div)
  })
}
loadCompanies()

