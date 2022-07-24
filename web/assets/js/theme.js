(function () {
    if ($(window).width() > 600) {
        return;
    }

    var $book = $(".book");
    var $summary = $('.book-summary');

    setTimeout(function () {
        var $toggleSummary = $('.toggle-summary');
        $toggleSummary.on('click', function () {
            var summaryOffset = null;
            var bookBodyOffset = null;
            if (isOpen()) {
                summaryOffset = -($summary.outerWidth());
                bookBodyOffset = 0;
                $book.removeClass('with-summary');
            } else {
                summaryOffset = 0;
                bookBodyOffset = $summary.outerWidth();
                $book.addClass('with-summary');
            }
        });
    }, 1);

    function isOpen() {
        return $book.hasClass("with-summary");
    }
})();