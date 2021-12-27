import {disableForm, enableForm} from "../forms.js";

document.getElementById("post-form").addEventListener("submit", function(e) {
  e.preventDefault();

  const imageFile = document.querySelector("#image-file").files[0];
  if (!imageFile || imageFile.type !== "image/jpeg") {
    alert("Please provide a .jpg file.");
    return;
  }
  const postBody = document.getElementById("body").value;
  if (postBody.length > 4096) {
    alert("Post body too long (max 4096 characters)");
    return;
  }
  const imageFileName = `${imageFile.name}?${Math.random().toString().replace("0.", "")}`;

  disableForm();

  getSignedUploadURL(imageFile, imageFileName)
    .then(r => r.json())
    .then(data => uploadFile(data.url, imageFile))
    .then(() => savePost(imageFileName, postBody))
    .then(r => r.json())
    .then(data => pollPost(data.id))
    .catch(err => {
      alert("Error saving post");
      console.log(err);
      enableForm();
    });
});

function getSignedUploadURL(imageFile, imageFileName) {
  return fetch("/api/signed_upload_url", {
    method: "POST",
    headers: new Headers({"Content-Type": "application/json"}),
    body: JSON.stringify({
      fileName: imageFileName,
      contentType: imageFile.type,
    }),
  });
}

function uploadFile(url, imageFile) {
  return fetch(url, {
    method: "PUT",
    headers: new Headers({"Content-Type": imageFile.type}),
    body: imageFile,
  });
}

function savePost(imageFileName, postBody) {
  return fetch("/api/create_post", {
    method: "POST",
    headers: new Headers({"Content-Type": "application/json"}),
    body: JSON.stringify({
      body: postBody,
      images: [imageFileName],
    }),
  });
}

function pollPost(postID) {
  fetch(`/api/post/${postID}`)
    .then(r => r.json())
    .then(data => {
      if (data.publishedAt) {
        location.href = "/";
      } else {
        setTimeout(pollPost, 1000, postID);
      }
    })
    .catch(err => {
      alert("Error fetching the created post.");
      console.log(err);
    });
}
