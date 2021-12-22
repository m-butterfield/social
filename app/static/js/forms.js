export function submitForm(e) {
  e.preventDefault();
  fetch(e.target.action, {
    method: "POST",
    headers: new Headers({"Content-Type": "application/json"}),
    body: JSON.stringify(Object.fromEntries(new FormData(e.target).entries())),
  }).then(resp => {
    if (resp.status === 400) {
      resp.text().then(val => alert(val));
    } else if (resp.status === 200 || resp.status === 201) {
      location.href = "/";
    }
  }).catch(err => {
    alert("error...");
    console.log(err);
  });
}

export function disableForm() {
  document.querySelectorAll(".upload-form-element").forEach(e => e.disabled = true);
}

export function enableForm() {
  document.querySelectorAll(".upload-form-element").forEach(e => e.disabled = false);
}
