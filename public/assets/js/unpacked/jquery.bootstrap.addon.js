/* ============================ TOC ==============================
1. Bootstrap Tab Collapse 
2. Initialize Scrollspy 
3. Sticky Sidebar for Scrollspy Content (is not part of Scrollspy) 
4. .nav-tabs 
5. Tab Collapse = if the tabs are in a sidebar/widget column 
6. Tab Collapse = if the tabs are in a narrow column 
7. Tab Collapse = if the tabs are in a wide-ish column 
8 a. ToolTips and Popover on hover 
8 b. ToolTips and Popover on tap/click 
9. Toggle new pagination 
10. .theme-carousel 
11. Enable swiping on carousel requires touch-swipe.js
12. Hide Tooltip when clicked off of click another tooltip
13. Hide Popover when clicked off of click another popover
================================================================ */
// --------------- 1. Bootstrap Tab Collapse --------------- //
/*!
 * Small bootstrap plugin that switches bootstrap tabs component to collapse component for small screens.
 * The MIT License (MIT) // Copyright (c) 2014 flatlogic.com // https://github.com/flatlogic/bootstrap-tabcollapse
 */
! function(b) {
    var a = function(d, c) {
        this.options = c;
        this.$tabs = b(d);
        this._accordionVisible = false;
        this._initAccordion();
        this._checkStateOnResize();
        this.checkState()
    };
    a.DEFAULTS = {
        accordionClass: "visible-xs",
        tabsClass: "hidden-xs",
        accordionTemplate: function(e, c, f, d) {
            return '<div class="panel panel-default">   <div class="panel-heading">      <h4 class="panel-title">        <a class="' + (d ? "" : "collapsed") + '" data-toggle="collapse" data-parent="#' + f + '" href="#' + c + '">           ' + e + '        </a>      </h4>   </div>   <div id="' + c + '" class="panel-collapse collapse ' + (d ? "in" : "") + '">       <div class="panel-body">       </div>   </div></div>'
        }
    };
    a.prototype.checkState = function() {
        if (this.$tabs.is(":visible") && this._accordionVisible) {
            this.showTabs();
            this._accordionVisible = false
        } else {
            if (this.$accordion.is(":visible") && !this._accordionVisible) {
                this.showAccordion();
                this._accordionVisible = true
            }
        }
    };
    a.prototype.showTabs = function() {
        this.$tabs.trigger(b.Event("show-tabs.bs.tabcollapse"));
        var c = this.$accordion.find(".panel-body");
        c.each(function() {
            var e = b(this),
                d = e.data("bs.tabcollapse.tabpane");
            d.append(e.children("*").detach())
        });
        this.$accordion.html("");
        this.$tabs.trigger(b.Event("shown-tabs.bs.tabcollapse"))
    };
    a.prototype.showAccordion = function() {
        this.$tabs.trigger(b.Event("show-accordion.bs.tabcollapse"));
        var d = this.$tabs.find('li:not(.dropdown) [data-toggle="tab"], li:not(.dropdown) [data-toggle="pill"]'),
            c = this;
        d.each(function() {
            var e = b(this);
            c.$accordion.append(c._createAccordionGroup(c.$accordion.attr("id"), e))
        });
        this.$tabs.trigger(b.Event("shown-accordion.bs.tabcollapse"))
    };
    a.prototype._checkStateOnResize = function() {
        var c = this;
        b(window).resize(function() {
            clearTimeout(c._resizeTimeout);
            c._resizeTimeout = setTimeout(function() {
                c.checkState()
            }, 100)
        })
    };
    a.prototype._initAccordion = function() {
        this.$accordion = b('<div class="panel-group ' + this.options.accordionClass + '" id="' + this.$tabs.attr("id") + '-accordion"></div>');
        this.$tabs.after(this.$accordion);
        this.$tabs.addClass(this.options.tabsClass);
        this.$tabs.siblings(".tab-content").addClass(this.options.tabsClass)
    };
    a.prototype._createAccordionGroup = function(i, f) {
        var e = f.attr("data-target"),
            g = f.parent().is(".active");
        if (!e) {
            e = f.attr("href");
            e = e && e.replace(/.*(?=#[^\s]*$)/, "")
        }
        var c = b(e),
            d = c.attr("id") + "-collapse",
            h = b(this.options.accordionTemplate(f.html(), d, i, g));
        h.find(".panel-body").append(c.children("*").detach()).data("bs.tabcollapse.tabpane", c);
        return h
    };
    b.fn.tabCollapse = function(c) {
        return this.each(function() {
            var f = b(this);
            var e = f.data("bs.tabcollapse");
            var d = b.extend({}, a.DEFAULTS, f.data(), typeof c == "object" && c);
            if (!e) {
                f.data("bs.tabcollapse", new a(this, d))
            }
        })
    };
    b.fn.tabCollapse.Constructor = a
}(window.jQuery);


// --------------- 2. Initialize Scrollspy --------------- //
$(document).ready(function() {
    var $window = $(window);
    var $body = $(document.body);
    $body.scrollspy({
        target: ".spying-sidebar",
        offset: 160 // == header plus the height of the headline and some margin
    });
    $window.on("load", function() {
        $body.scrollspy("refresh");
    });
    $(".affix-page-container .spying-sidebar [href=#]").click(function(e) {
        e.preventDefault();
    });
});


// --------------- 3. Sticky Sidebar for Scrollspy Content (is not part of Scrollspy) --------------- //
    if ($(".spying-sidebar").length > 0) {
        var $content = $(".primary-column"),
            $sidebar = $(".spying-sidebar"),
            $window = $(window), 
            $document = $(document);

        function stickyScrollSpySideBar() {
            var position = $('#sidebar').position().left + 15;
            var headerHeight = $('.header').outerHeight();
            var colwidth = $sidebar.outerWidth(true);
            $sidebar.css('width', (colwidth) + 'px');
            $sidebar.css('top', (headerHeight) + 'px');
            $sidebar.css('left', (position) + 'px');

            var fixedBottom = $document.height() - ($content.offset().top + $content.height());
            var topOffset = $content.offset().top;
            var stopPoint = $document.height() - (fixedBottom + $sidebar.height() + 40);
            if ($content.height() > $sidebar.height() + 0) {
                if ($window.scrollTop() > topOffset) {
                    $sidebar.addClass("sticky-spy-bar");
                    if ($window.scrollTop() > stopPoint) {
                        $sidebar.css("top", (stopPoint - $window.scrollTop() + 40));
                    }
                } else {
                    $sidebar.removeClass("sticky-spy-bar");
                    $sidebar.removeAttr('style');
                }
            }
        }
        stickyScrollSpySideBar();
        window.onscroll = function() {
            stickyScrollSpySideBar();
        };
        window.onresize = function() {
            stickyScrollSpySideBar();
        };
    }

$(document).ready(function() {

    // --------------- 4. .nav-tabs --------------- //
    $('.nav-tabs a').click(function(e) {
        e.preventDefault();
        $(this).tab('show');
    });

    // --------------- 5. Tab Collapse = if the tabs are in a sidebar/widget column --------------- //
    $('.sidebar-tabs').tabCollapse({
        tabsClass: 'sidebar-tabcollapse-visible-xl', // roll your own custom classes DO NOT USE responsive utility classes
        accordionClass: 'sidebar-tabcollapse-hidden-xl' // roll your own custom classes DO NOT USE responsive utility classes
    });

    // --------------- 6. Tab Collapse = if the tabs are in a narrow column --------------- //
    $('.narrow-tabs').tabCollapse({
        tabsClass: 'narrow-tabcollapse-visible-xl', // roll your own custom classes DO NOT USE responsive utility classes
        accordionClass: 'narrow-tabcollapse-hidden-xl' // roll your own custom classes DO NOT USE responsive utility classes
    });

    // --------------- 7. Tab Collapse = if the tabs are in a wide-ish column --------------- //
    $('.content-tabs').tabCollapse({
        tabsClass: 'content-tabcollapse-visible-xl', // roll your own custom classes DO NOT USE responsive utility classes
        accordionClass: 'content-tabcollapse-hidden-xl' // roll your own custom classes DO NOT USE responsive utility classes
    });

    // ---------------  8 a. ToolTips and Popover on hover --------------- //
    if ($('html').hasClass('no-touch')) {
        $('[data-hover="tooltip"]').tooltip({
            trigger: 'hover',
            container: 'body'
        });
        $('[data-hover="popover"]').popover({
            trigger: 'hover'
        });
    }

    // --------------- 8 b. ToolTips and Popover on tap/click --------------- //
    $('[data-toggle="tooltip"]').tooltip({
        trigger: 'click',
        container: 'body'
    });
    $('[data-toggle="popover"]').popover({
        trigger: 'click'
    });

    // --------------- 9. Toggle new pagination --------------- //
    $('.new-pagination-toggle').click(function(f) {
        $(this).next('.new-pagination').slideToggle();
        $(this).toggleClass('active');
        f.preventDefault()
    });

    // ---------------  10. .theme-carousel --------------- //
    $('.theme-carousel').carousel({
        //interval: 3000
    }).on('slide.bs.carousel', function(e) {
        var nextH = $(e.relatedTarget).height();
        $(this).find('div.active').parent().animate({
            height: nextH
        }, 500);
    });


    // --------------- 11. Enable swiping on carousel requires touch-swipe.js --------- //
    $(".carousel-inner").swipe({
        swipeLeft: function(event, direction, distance, duration, fingerCount) {
            $(this).parent().carousel('prev');
        },
        swipeRight: function() {
            $(this).parent().carousel('next');
        },
        threshold: 50
    });

    $(".carousel-inner").on({
        "mousedown": function(e) {
            $(this).addClass("mouseDown");
            return false;
        },
        "mouseup": function() {
            $(this).removeClass("mouseDown");
        }
    });

});


// --------------- 12. Hide Tooltip when clicked off of click another tooltip --------- //
$('body').on('mouseup', function(e) {
    $('[data-toggle="tooltip"]').each(function() {
        if (!$(this).is(e.target) && $(this).has(e.target).length === 0 && $('.tooltip').has(e.target).length === 0) {
            $(this).tooltip('hide');
        }
    });
});

// --------------- 13. Hide Popover when clicked off of click another popover ===== //
$('body').on('mouseup', function(e) {
    $('[data-toggle="popover"]').each(function() {
        if (!$(this).is(e.target) && $(this).has(e.target).length === 0 && $('.popover').has(e.target).length === 0) {
            $(this).popover('hide');
        }
    });
});