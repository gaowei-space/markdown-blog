(function () {
    // MEMO:
    // Gitbook is calculated as "calc (100% - 60px)" in the horizontal width when the width of the screen size is 600px
    // or less.
    // In this case, since contradiction occurs in the implementation of this module, return.
    if ($(window).width() <= 600) {
        return;
    }

    var KEY_SPLIT_STATE = 'plugin_gitbook_split';

    var dividerWidth = null;
    var isDraggable = false;
    var dividerCenterOffsetLeft = null;
    var splitState = null;
    var grabPointWidth = null;

    var $body = $('body');
    var $book = $('.book');
    var $summary = $('.book-summary');
    var $bookBody = $('.book-body');
    var $divider = $('<div class="divider-content-summary">' +
        '<div class="divider-content-summary__icon">' +
        '<i class="fa fa-ellipsis-v"></i>' +
        '</div>' +
        '</div>');

    $summary.append($divider);

    dividerWidth = $divider.outerWidth();
    dividerCenterOffsetLeft = $divider.outerWidth() / 2;

    // restore split state from sessionStorage
    splitState = getSplitState();

    // 根据缓存来判断当前页面展开的状态
    // 因为刷新之后，with-summary默认存在，这里需要根据状态做调整
    if (splitState.bookBodyOffset === 0) {
        $book.removeClass('with-summary');
    }

    setSplitState(
        splitState.summaryWidth,
        splitState.summaryOffset,
        splitState.bookBodyOffset
    );

    setTimeout(function () {
        var $toggleSummary = $('.toggle-summary');

        $toggleSummary.on('click', function () {

            var summaryOffset = null;
            var bookBodyOffset = null;

            var isOpen = $book.hasClass('with-summary');

            if (isOpen) {
                summaryOffset = -($summary.outerWidth());
                bookBodyOffset = 0;
                $book.removeClass('with-summary');
            } else {
                summaryOffset = 0;
                bookBodyOffset = $summary.outerWidth();
                $book.addClass('with-summary');
            }

            setSplitState($summary.outerWidth(), summaryOffset, bookBodyOffset);
            saveSplitState($summary.outerWidth(), summaryOffset, bookBodyOffset);
        });
    }, 1);

    $divider.on('mousedown', function (event) {
        event.stopPropagation();
        isDraggable = true;
        grabPointWidth = $summary.outerWidth() - event.pageX;
    });

    $body.on('mouseup', function (event) {
        event.stopPropagation();
        isDraggable = false;
        saveSplitState(
            $summary.outerWidth(),
            $summary.position().left,
            $bookBody.position().left
        );
    });

    $body.on('mousemove', function (event) {
        if (!isDraggable) {
            return;
        }
        event.stopPropagation();
        event.preventDefault();
        $summary.outerWidth(event.pageX + grabPointWidth);
        $bookBody.offset({ left: event.pageX + grabPointWidth });
    });

    function getSplitState() {
        var splitState = JSON.parse(sessionStorage.getItem(KEY_SPLIT_STATE)) || {};
        splitState.summaryWidth !== undefined || (splitState.summaryWidth = $summary.outerWidth());
        splitState.summaryOffset !== undefined || (splitState.summaryOffset = $summary.position().left);
        splitState.bookBodyOffset !== undefined || (splitState.bookBodyOffset = $bookBody.position().left);
        return splitState;
    }

    function saveSplitState(summaryWidth, summaryWidthOffset, bookBodyOffset) {
        sessionStorage.setItem(KEY_SPLIT_STATE, JSON.stringify({
            summaryWidth: summaryWidth,
            summaryOffset: summaryWidthOffset,
            bookBodyOffset: bookBodyOffset,
        }));
    }

    function setSplitState(summaryWidth, summaryOffset, bookBodyOffset) {
        $summary.outerWidth(summaryWidth);
        $summary.offset({ left: summaryOffset });
        $bookBody.offset({ left: bookBodyOffset });
        // improved broken layout in windows chrome.
        //   "$(x).offset" automatically add to "position:relative".
        //   but it cause layout broken..
        $summary.css({ position: 'absolute' });
        $bookBody.css({ position: 'absolute' });
    }
})();