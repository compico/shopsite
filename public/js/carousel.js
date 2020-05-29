bulmaCarousel.attach('#carousel-pens', {
    slidesToScroll: 1,
    slidesToShow: 1,
    infinite: true
});
var element = document.querySelector('#carousel-pens');
if (element && element.bulmaCarousel) {
    // bulmaCarousel instance is available as element.bulmaCarousel
}