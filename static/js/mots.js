/*globals jQuery, document */
(function ($) {
    "use strict";

    $(document).ready(function(){

        /* Responsive videos */
        $(".post").fitVids();

        /* Responsive typography */
        $(".blog-logo").fitText(0.8, { minFontSize: '36px', maxFontSize: '54px' });
        $(".post-title").fitText(0.8, { minFontSize: '36px', maxFontSize: '54px' });
    });

}(jQuery));