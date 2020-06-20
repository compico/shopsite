let elstars = [];
let hiddenvar = document.getElementById("stars");
let clicked = false;
let debug = [];
document.addEventListener('DOMContentLoaded', () => {
    for (let i = 1; i < 6; i++) {
        x = document.getElementsByClassName("star" + i)[0];
        x.position = i;
        elstars.push(x);
    }
    elstars.forEach(el => {
        el.addEventListener("mouseover", () => {
            if (!clicked) {
                for (let i = el.position - 1; i > -1; i--) {
                    elstars[i].lastChild.src = "/public/image/svg/star.svg";
                }
            }
        });
        el.addEventListener("mouseleave", () => {
            if (!clicked) {
                for (let i = el.position - 1; i > -1; i--) {
                    elstars[i].lastChild.src = "/public/image/svg/star2.svg";
                }
            }
        });
        var debvar = 0;
        el.addEventListener("click", () => {
            for (let i = el.position - 1; i > -1; i--) {
                elstars[i].lastChild.src = "/public/image/svg/star.svg";
            }
            for (let i = el.position; i < 5; i++) {
                elstars[i].lastChild.src = "/public/image/svg/star2.svg";
            }
            clicked = true;
            hiddenvar.value = el.position;
        });
    });
});