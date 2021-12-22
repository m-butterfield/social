import {disableForm, enableForm} from "../forms.js";

document.getElementById("post-form").addEventListener("submit", async function(e) {
  e.preventDefault();
  const imageFile = document.querySelector("#image-file").files[0];
  if (!imageFile || imageFile.type !== "image/jpeg") {
    alert("Please provide a .jpg file.");
    return;
  }

  disableForm();

  const imageFileName = `${imageFile.name}?${Math.random().toString().replace("0.", "")}`;

  await uploadFile(imageFile, imageFileName).catch(err => {
    alert("error uploading image");
    console.log(err);
  });

  console.log("saving post...");
  enableForm();
});

function uploadFile(fileObj, fileName) {
  return fetch("/user/signed_upload_url", {
    method: "POST",
    headers: new Headers({"Content-Type": "application/json"}),
    body: JSON.stringify({
      fileName: fileName,
      contentType: fileObj.type,
    }),
  }).then(r => r.json()).then(data => {
    return upload(data.url, fileObj);
  });
}

function upload(url, file) {
  return fetch(url, {
    method: "PUT",
    headers: new Headers({
      "Content-Type": file.type,
    }),
    body: file,
  });
}
