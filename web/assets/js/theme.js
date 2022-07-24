(function () {
    if ($(window).width() > 600) {
        return;
    }

    var $book = $(".book");
    setTimeout(function () {
        $('.toggle-summary').on('click', function () {
            $book.toggleClass('with-summary', !$book.hasClass("with-summary"));
        });
    }, 1);
})();