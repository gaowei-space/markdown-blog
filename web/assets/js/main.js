(function () {
    hljs.highlightAll();

    var KEY_THEME_STATE = 'blog_theme_state';
    var $book = $('.book');
    var $themeAction = $('.js-theme-action');
    var $themeCss = document.getElementById('theme-css');

    changeTheme(true)

    $themeAction.on('click', function (event) {
        event.stopPropagation();
        changeTheme(false)
    });

    function changeTheme(isInit = false) {
        color = isInit ? getThemeState().color : (getThemeState().color == 'dark' ? 'white' : 'dark')

        setThemeState(color)
        saveThemeState(color)
    }

    function getThemeState() {
        var themeState = JSON.parse(sessionStorage.getItem(KEY_THEME_STATE)) || {};
        themeState.color !== undefined || (themeState.color = 'dark');
        return themeState;
    }

    function saveThemeState(color) {
        sessionStorage.setItem(KEY_THEME_STATE, JSON.stringify({
            color: color,
        }));
    }

    function setThemeState(color) {
        if (color == 'dark') {
            $book.addClass('color-theme-2');
            $themeAction.html('<i class="fa fa-sun-o"></i>')
            $themeCss.href = "/static/css/github-markdown-css/dark.css"
        } else {
            $book.removeClass('color-theme-2');
            $themeAction.html('<i class="fa fa-moon-o"></i>')
            $themeCss.href = "/static/css/github-markdown-css/white.css"
        }
    }
})();