import { submitForm } from "./forms.js";

Array.from(document.getElementsByClassName("user-form")).forEach(f => {
  f.addEventListener("submit", submitForm);
});
