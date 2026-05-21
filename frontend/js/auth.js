async function register() {

  const email =
    document.getElementById("registerEmail").value

  const password =
    document.getElementById("registerPassword").value

  const response = await fetch(
    API + "/register",
    {
      method: "POST",

      headers: {
        "Content-Type": "application/json"
      },

      body: JSON.stringify({
        email,
        password
      })
    }
  )

  const data = await response.text()

  show(data)
}

async function login() {

  const email =
    document.getElementById("loginEmail").value

  const password =
    document.getElementById("loginPassword").value

  const response = await fetch(
    API + "/login",
    {
      method: "POST",

      headers: {
        "Content-Type": "application/json"
      },

      body: JSON.stringify({
        email,
        password
      })
    }
  )

  const data = await response.json()

  
  localStorage.setItem(
    "token",
    data.token,
  )

  
  window.location.href = "dashboard.html"
}

function show(text) {
  document.getElementById("output")
    .textContent = text
}