let elstars = [];
let hiddenvar = document.getElementById("stars");
let clicked = false;
let debug = [];
let subbtm = document.getElementById("submit-button");
let author = document.querySelector("input[name=author]");
let stars = document.querySelector("input[name=stars]");
let reviewtext = document.querySelector("textarea[name=reviewtext]");
function check() {
  if (author.value != "" && stars.value != 0 && reviewtext.value != "") {
    if (author.value.length > 6 && reviewtext.value.length > 16) {
      subbtm.disabled = false;
      return;
    }
  }
  subbtm.disabled = true;
}
document.addEventListener('DOMContentLoaded', () => {
  for (let i = 0; i < 5; i++) {
    x = document.getElementsByClassName("star" + (i + 1))[0];
    x.position = i;
    elstars.push(x);
  }
  author.addEventListener("input", () => {
    check();
  });
  stars.addEventListener("input", () => {
    check();
  });
  reviewtext.addEventListener("input", () => {
    check();
  });
  elstars.forEach(el => {
    el.addEventListener("mouseover", () => {
      if (!clicked) {
        for (let i = el.position; i >= 0; i--) {
          elstars[i].lastChild.src = "/public/image/svg/star2.svg";
        }
      }
      el.style.cssText =
        `-webkit-filter: drop-shadow( 0px 0px 5px rgb(255, 193, 7));
      filter: drop-shadow( 0px 0px 5px rgb(255, 193, 7));`;
    });
    el.addEventListener("mouseleave", () => {
      if (!clicked) {
        for (let i = el.position; i >= 0; i--) {
          elstars[i].lastChild.src = "/public/image/svg/star1.svg";
        }
      }
      el.style.cssText = "";
    });
    var debvar = 0;
    el.addEventListener("click", () => {
      for (let i = el.position; i >= 0; i--) {
        elstars[i].lastChild.src = "/public/image/svg/star2.svg";
      }
      for (let i = el.position + 1; i < 5; i++) {
        elstars[i].lastChild.src = "/public/image/svg/star1.svg";
      }
      clicked = true;
      hiddenvar.value = el.position + 1;
      check();
    });
  });
});

disabelel = (el) => { el.lastChild.src = "/public/image/svg/star1.svg"; };
enableel = (el) => { el.lastChild.src = "/public/image/svg/star2.svg"; };