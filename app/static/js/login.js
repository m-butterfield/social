function submitUser(e) {
  e.preventDefault();
  fetch(e.target.action, {
    method: "POST",
    headers: new Headers({"Content-Type": "application/json"}),
    body: JSON.stringify(Object.fromEntries(new FormData(e.target).entries())),
  }).then(resp => {
    if (resp.status === 400) {
      resp.text().then(val => alert(val));
    } else if (resp.status === 200) {
      alert("success!");
      location.href = "/";
    }
  }).catch(err => {
    alert("error...");
    console.log(err);
  });
}

Array.from(document.getElementsByClassName("user-form")).forEach(f => {
  f.addEventListener("submit", submitUser);
});
