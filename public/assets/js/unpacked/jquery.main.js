/* ============================ TOC ==============================
1.  IOS 8 Bug Fix
2.  BIGTEXT
3.  JQUERY COOKIE
4.  HOVER INTENT ON THE .hover function in jQuery
5.  SUPPORTS TOUCH OR NOT for IOS, Android, and Windows Mobile
6.  RESPONSIVE EQUAL HEIGHTS
7.  FITVIDEOS FOR RESPONSIVE IFRAMES FROM YOUTUBE, VIMEO, SOUND CLOUD ETC.,
8.  PLACEHOLDER SUPPORT
9.  MENU SCRIPT
10. INITIALIZE MENU
11. SOUND CLOUD NON-BLOCKING LOAD
12. BIG TEXT
13. DATA SLIDE
14. DON'T COVER ANCHOR TARGETS
15. STICKY HEADER
16. MATCH HEIGHT OF SLIDER NEXT TO THIS COLUMN
17. MISC INITIALIZATIONS AND CUSTOM SCRIPTS
    17. - 1 :: Equal Heights Initialize
    17. - 2 :: .title inner wrapper
    17. - 3 :: Inner Wrappers for .info-text-box and .widget-title
    17. - 4 :: icon-click functions :: you can remove this on live sites
    17. - 5 :: Scroll to Top
    17. - 6 :: Video Non-Blocking Load
    17. - 7 :: pricing plan
    17. - 8 :: .click-slide
18. STICKY ADD THIS AND SIDEBAR FOLLOW
19. SUMMARY PAGE SIDEBAR FOLLOW
20. JQUERY STICKY FOOTER
================================================================ */


// --------------- 1. IOS 8 Bug Fix --------------- //
$(window).load(function() {
    var deviceAgent = navigator.userAgent.toLowerCase();
    var agentID = deviceAgent.match(/(iphone|ipod|ipad)/);
    if (agentID) {
        setTimeout(function() {
            window.scrollTo(1, 1);
        }, 100);
    }
});

/* --------------- 2. BIGTEXT --------------- */
/*! BigText - v0.1.7a - 2014-07-18
 * https://github.com/zachleat/bigtext
 * Copyright (c) 2014 Zach Leatherman (@zachleat)
 * MIT License */

(function(g, e) {
    var a = 0,
        i = e("head"),
        h = g.BigText,
        f = e.fn.bigtext,
        d = {
            DEBUG_MODE: false,
            DEFAULT_MIN_FONT_SIZE_PX: null,
            DEFAULT_MAX_FONT_SIZE_PX: 528,
            GLOBAL_STYLE_ID: "bigtext-style",
            STYLE_ID: "bigtext-id",
            LINE_CLASS_PREFIX: "bigtext-line",
            EXEMPT_CLASS: "bigtext-exempt",
            noConflict: function(j) {
                if (j) {
                    e.fn.bigtext = f;
                    g.BigText = h
                }
                return d
            },
            test: {
                wholeNumberFontSizeOnly: function() {
                    if (!("getComputedStyle" in g) || document.body == null) {
                        return true
                    }
                    var k = e("<div/>").css({
                            position: "absolute",
                            "font-size": "14.1px"
                        }).appendTo(document.body).get(0),
                        j = g.getComputedStyle(k, null);
                    if (j) {
                        return j.getPropertyValue("font-size") === "14px"
                    }
                    return true
                }
            },
            supports: {
                wholeNumberFontSizeOnly: undefined
            },
            init: function() {
                if (d.supports.wholeNumberFontSizeOnly === undefined) {
                    d.supports.wholeNumberFontSizeOnly = d.test.wholeNumberFontSizeOnly()
                }
                if (!e("#" + d.GLOBAL_STYLE_ID).length) {
                    i.append(d.generateStyleTag(d.GLOBAL_STYLE_ID, [".bigtext * { white-space: nowrap; } .bigtext > * { display: block; }", ".bigtext ." + d.EXEMPT_CLASS + ", .bigtext ." + d.EXEMPT_CLASS + " * { white-space: normal; }"]))
                }
            },
            bindResize: function(j, k) {
                if (e.throttle) {
                    e(g).unbind(j).bind(j, e.throttle(100, k))
                } else {
                    if (e.fn.smartresize) {
                        j = "smartresize." + j
                    }
                    e(g).unbind(j).bind(j, k)
                }
            },
            getStyleId: function(j) {
                return d.STYLE_ID + "-" + j
            },
            generateStyleTag: function(k, j) {
                return e("<style>" + j.join("\n") + "</style>").attr("id", k)
            },
            clearCss: function(k) {
                var j = d.getStyleId(k);
                e("#" + j).remove()
            },
            generateCss: function(r, q, p, o) {
                var n = [];
                d.clearCss(r);
                for (var m = 0, l = q.length; m < l; m++) {
                    n.push("#" + r + " ." + d.LINE_CLASS_PREFIX + m + " {" + (o[m] ? " white-space: normal;" : "") + (q[m] ? " font-size: " + q[m] + "px;" : "") + (p[m] ? " word-spacing: " + p[m] + "px;" : "") + "}")
                }
                return d.generateStyleTag(d.getStyleId(r), n)
            },
            jQueryMethod: function(j) {
                d.init();
                j = e.extend({
                    minfontsize: d.DEFAULT_MIN_FONT_SIZE_PX,
                    maxfontsize: d.DEFAULT_MAX_FONT_SIZE_PX,
                    childSelector: "",
                    resize: true
                }, j || {});
                this.each(function() {
                    var o = e(this).addClass("bigtext"),
                        m = o.width(),
                        n = o.attr("id"),
                        k = j.childSelector ? o.find(j.childSelector) : o.children();
                    if (!n) {
                        n = "bigtext-id" + (a++);
                        o.attr("id", n)
                    }
                    if (j.resize) {
                        d.bindResize("resize.bigtext-event-" + n, function() {
                            d.jQueryMethod.call(e("#" + n), j)
                        })
                    }
                    d.clearCss(n);
                    k.addClass(function(p, q) {
                        return [q.replace(new RegExp("\\b" + d.LINE_CLASS_PREFIX + "\\d+\\b"), ""), d.LINE_CLASS_PREFIX + p].join(" ")
                    });
                    var l = b(o, k, m, j.maxfontsize, j.minfontsize);
                    i.append(d.generateCss(n, l.fontSizes, l.wordSpacings, l.minFontSizes))
                });
                return this.trigger("bigtext:complete")
            }
        };

    function c(p, r, q, s, k, n, m) {
        var j;
        m = typeof m === "number" ? m : 0;
        p.css(q, s + n);
        j = p.width();
        if (j >= r) {
            p.css(q, "");
            if (j === r) {
                return {
                    match: "exact",
                    size: parseFloat((parseFloat(s) - 0.1).toFixed(3))
                }
            }
            var o = r - m,
                l = j - r;
            return {
                match: "estimate",
                size: parseFloat((parseFloat(s) - (q === "word-spacing" && m && (l < o) ? 0 : k)).toFixed(3))
            }
        }
        return j
    }

    function b(m, q, r, s, p) {
        var k = m.clone(true).addClass("bigtext-cloned").css({
            fontFamily: m.css("font-family"),
            textTransform: m.css("text-transform"),
            wordSpacing: m.css("word-spacing"),
            letterSpacing: m.css("letter-spacing"),
            position: "absolute",
            left: d.DEBUG_MODE ? 0 : -9999,
            top: d.DEBUG_MODE ? 0 : -9999
        }).appendTo(document.body);
        var l = [],
            j = [],
            o = [],
            n = [];
        q.css("float", "left").each(function() {
            var C = e(this),
                x = d.supports.wholeNumberFontSizeOnly ? [8, 4, 1] : [8, 4, 1, 0.1],
                z, B;
            if (C.hasClass(d.EXEMPT_CLASS)) {
                l.push(null);
                n.push(null);
                o.push(false);
                return
            }
            var D = 32,
                A = parseFloat(C.css("font-size")),
                y = (C.width() / A).toFixed(6);
            B = parseInt(r / y, 10) - D;
            outer: for (var u = 0, t = x.length; u < t; u++) {
                inner: for (var w = 1, v = 10; w <= v; w++) {
                    if (B + w * x[u] > s) {
                        B = s;
                        break outer
                    }
                    z = c(C, r, "font-size", B + w * x[u], x[u], "px", z);
                    if (typeof z !== "number") {
                        B = z.size;
                        if (z.match === "exact") {
                            break outer
                        }
                        break inner
                    }
                }
            }
            n.push(r / B);
            if (B > s) {
                l.push(s);
                o.push(false)
            } else {
                if (!!p && B < p) {
                    l.push(p);
                    o.push(true)
                } else {
                    l.push(B);
                    o.push(false)
                }
            }
        }).each(function(u) {
            var y = e(this),
                x = 0,
                v = 1,
                w;
            if (y.hasClass(d.EXEMPT_CLASS)) {
                j.push(null);
                return
            }
            y.css("font-size", l[u] + "px");
            for (var t = 1, z = 3; t < z; t += v) {
                w = c(y, r, "word-spacing", t, v, "px", w);
                if (typeof w !== "number") {
                    x = w.size;
                    break
                }
            }
            y.css("font-size", "");
            j.push(x)
        }).removeAttr("style");
        if (!d.DEBUG_MODE) {
            k.remove()
        } else {
            k.css({
                "background-color": "rgba(255,255,255,.4)"
            })
        }
        return {
            fontSizes: l,
            wordSpacings: j,
            ratios: n,
            minFontSizes: o
        }
    }
    e.fn.bigtext = d.jQueryMethod;
    g.BigText = d
})(this, jQuery);


/* --------------- 3. JQUERY COOKIE --------------- */
/*!jQuery Cookie Plugin v1.4.1 :  https://github.com/carhartl/jquery-cookie
 * Copyright 2006, 2014 Klaus Hartl : Released under the MIT license */
(function(a) {
    if (typeof define === "function" && define.amd) {
        define(["jquery"], a)
    } else {
        if (typeof exports === "object") {
            a(require("jquery"))
        } else {
            a(jQuery)
        }
    }
}(function(f) {
    var a = /\+/g;

    function d(i) {
        return b.raw ? i : encodeURIComponent(i)
    }

    function g(i) {
        return b.raw ? i : decodeURIComponent(i)
    }

    function h(i) {
        return d(b.json ? JSON.stringify(i) : String(i))
    }

    function c(i) {
        if (i.indexOf('"') === 0) {
            i = i.slice(1, -1).replace(/\\"/g, '"').replace(/\\\\/g, "\\")
        }
        try {
            i = decodeURIComponent(i.replace(a, " "));
            return b.json ? JSON.parse(i) : i
        } catch (j) {}
    }

    function e(j, i) {
        var k = b.raw ? j : c(j);
        return f.isFunction(i) ? i(k) : k
    }
    var b = f.cookie = function(q, p, v) {
        if (arguments.length > 1 && !f.isFunction(p)) {
            v = f.extend({}, b.defaults, v);
            if (typeof v.expires === "number") {
                var r = v.expires,
                    u = v.expires = new Date();
                u.setTime(+u + r * 86400000)
            }
            return (document.cookie = [d(q), "=", h(p), v.expires ? "; expires=" + v.expires.toUTCString() : "", v.path ? "; path=" + v.path : "", v.domain ? "; domain=" + v.domain : "", v.secure ? "; secure" : ""].join(""))
        }
        var w = q ? undefined : {};
        var s = document.cookie ? document.cookie.split("; ") : [];
        for (var o = 0, m = s.length; o < m; o++) {
            var n = s[o].split("=");
            var j = g(n.shift());
            var k = n.join("=");
            if (q && q === j) {
                w = e(k, p);
                break
            }
            if (!q && (k = e(k)) !== undefined) {
                w[j] = k
            }
        }
        return w
    };
    b.defaults = {};
    f.removeCookie = function(j, i) {
        if (f.cookie(j) === undefined) {
            return false
        }
        f.cookie(j, "", f.extend({}, i, {
            expires: -1
        }));
        return !f.cookie(j)
    }
}));


/* --------------- 4. HOVER INTENT ON THE .hover function in jQuery ---------------*/
/*! hoverIntent v1.8.0 // 2014.06.29 // jQuery v1.9.1+
 * http://cherne.net/brian/resources/jquery.hoverIntent.html
 * You may use hoverIntent under the terms of the MIT license. Basically that
 * means you are free to use hoverIntent as long as this header is left intact.
 * Copyright 2007, 2014 Brian Cherne */
(function($) {
    $.fn.hoverIntent = function(handlerIn, handlerOut, selector) {
        var cfg = {
            interval: 100,
            sensitivity: 6,
            timeout: 3
        };
        if (typeof handlerIn === "object") {
            cfg = $.extend(cfg, handlerIn)
        } else {
            if ($.isFunction(handlerOut)) {
                cfg = $.extend(cfg, {
                    over: handlerIn,
                    out: handlerOut,
                    selector: selector
                })
            } else {
                cfg = $.extend(cfg, {
                    over: handlerIn,
                    out: handlerIn,
                    selector: handlerOut
                })
            }
        }
        var cX, cY, pX, pY;
        var track = function(ev) {
            cX = ev.pageX;
            cY = ev.pageY
        };
        var compare = function(ev, ob) {
            ob.hoverIntent_t = clearTimeout(ob.hoverIntent_t);
            if (Math.sqrt((pX - cX) * (pX - cX) + (pY - cY) * (pY - cY)) < cfg.sensitivity) {
                $(ob).off("mousemove.hoverIntent", track);
                ob.hoverIntent_s = true;
                return cfg.over.apply(ob, [ev])
            } else {
                pX = cX;
                pY = cY;
                ob.hoverIntent_t = setTimeout(function() {
                    compare(ev, ob)
                }, cfg.interval)
            }
        };
        var delay = function(ev, ob) {
            ob.hoverIntent_t = clearTimeout(ob.hoverIntent_t);
            ob.hoverIntent_s = false;
            return cfg.out.apply(ob, [ev])
        };
        var handleHover = function(e) {
            var ev = $.extend({}, e);
            var ob = this;
            if (ob.hoverIntent_t) {
                ob.hoverIntent_t = clearTimeout(ob.hoverIntent_t)
            }
            if (e.type === "mouseenter") {
                pX = ev.pageX;
                pY = ev.pageY;
                $(ob).on("mousemove.hoverIntent", track);
                if (!ob.hoverIntent_s) {
                    ob.hoverIntent_t = setTimeout(function() {
                        compare(ev, ob)
                    }, cfg.interval)
                }
            } else {
                $(ob).off("mousemove.hoverIntent", track);
                if (ob.hoverIntent_s) {
                    ob.hoverIntent_t = setTimeout(function() {
                        delay(ev, ob)
                    }, cfg.timeout)
                }
            }
        };
        return this.on({
            "mouseenter.hoverIntent": handleHover,
            "mouseleave.hoverIntent": handleHover
        }, cfg.selector)
    }
})(jQuery);


/* --------------- 5. SUPPORTS TOUCH OR NOT for IOS, Android, and Windows Mobile --------------- */
/*! Detects touch support and adds appropriate classes to html and returns a JS object  |  Copyright (c) 2013 Izilla Partners Pty Ltd  | http://www.izilla.com.au / Licensed under the MIT license  |  https://coderwall.com/p/egbgdw */
var supports = (function() {
    var d = document.documentElement,
        c = "ontouchstart" in window || navigator.msMaxTouchPoints;
    if (c) {
        d.className += " touch";
        return {
            touch: true
        }
    } else {
        d.className += " no-touch";
        return {
            touch: false
        }
    }
})();

/* --------------- 6. RESPONSIVE EQUAL HEIGHTS --------------- */
/* Javascript-Equal-Height-Responsive-Rows https://github.com/Sam152/Javascript-Equal-Height-Responsive-Rows */
(function($) {
    $.fn.equalHeight = function() {
        var heights = [];
        $.each(this, function(i, element) {
            $element = $(element);
            var element_height;
            var includePadding = ($element.css('box-sizing') == 'border-box') || ($element.css('-moz-box-sizing') == 'border-box');
            if (includePadding) {
                element_height = $element.innerHeight();
            } else {
                element_height = $element.height();
            }
            heights.push(element_height);
        });
        this.css('height', Math.max.apply(window, heights) + 'px');
        return this;
    };
    $.fn.equalHeightGrid = function(columns) {
        var $tiles = this;
        $tiles.css('height', 'auto');
        for (var i = 0; i < $tiles.length; i++) {
            if (i % columns === 0) {
                var row = $($tiles[i]);
                for (var n = 1; n < columns; n++) {
                    row = row.add($tiles[i + n]);
                }
                row.equalHeight();
            }
        }
        return this;
    };
    $.fn.detectGridColumns = function() {
        var offset = 0,
            cols = 0;
        this.each(function(i, elem) {
            var elem_offset = $(elem).offset().top;
            if (offset === 0 || elem_offset == offset) {
                cols++;
                offset = elem_offset;
            } else {
                return false;
            }
        });
        return cols;
    };
    $.fn.responsiveEqualHeightGrid = function() {
        var _this = this;

        function syncHeights() {
            var cols = _this.detectGridColumns();
            _this.equalHeightGrid(cols);
        }
        $(window).bind('resize load', syncHeights);
        syncHeights();
        return this;
    };
})(jQuery);


/* --------------- 7. FITVIDEOS FOR RESPONSIVE IFRAMES FROM YOUTUBE, VIMEO, SOUND CLOUD ETC., --------------- */
/*! FitVids 1.1 : Copyright 2013, Chris Coyier - http://css-tricks.com + Dave Rupert - http://daverupert.com
 * Credit to Thierry Koblentz - http://www.alistapart.com/articles/creating-intrinsic-ratios-for-video/
 * Released under the WTFPL license - http://sam.zoy.org/wtfpl/ */
(function(a) {
    a.fn.fitVids = function(b) {
        var e = {
            customSelector: null,
            ignore: null
        };
        if (!document.getElementById("fit-vids-style")) {
            var d = document.head || document.getElementsByTagName("head")[0];
            var c = ".fluid-width-video-wrapper{width:100%;position:relative;padding:0;}.fluid-width-video-wrapper iframe,.fluid-width-video-wrapper object,.fluid-width-video-wrapper embed {position:absolute;top:0;left:0;width:100%;height:100%;}";
            var f = document.createElement("div");
            f.innerHTML = '<p>x</p><style id="fit-vids-style">' + c + "</style>";
            d.appendChild(f.childNodes[1])
        }
        if (b) {
            a.extend(e, b)
        }
        return this.each(function() {
            var g = ["iframe[src*='player.vimeo.com']", "iframe[src*='youtube.com']", "iframe[src*='youtube-nocookie.com']", "iframe[src*='kickstarter.com'][src*='video.html']", "object", "embed"];
            if (e.customSelector) {
                g.push(e.customSelector)
            }
            var h = ".fitvidsignore";
            if (e.ignore) {
                h = h + ", " + e.ignore
            }
            var i = a(this).find(g.join(","));
            i = i.not("object object");
            i = i.not(h);
            i.each(function() {
                var n = a(this);
                if (n.parents(h).length > 0) {
                    return
                }
                if (this.tagName.toLowerCase() === "embed" && n.parent("object").length || n.parent(".fluid-width-video-wrapper").length) {
                    return
                }
                if ((!n.css("height") && !n.css("width")) && (isNaN(n.attr("height")) || isNaN(n.attr("width")))) {
                    n.attr("height", 9);
                    n.attr("width", 16)
                }
                var j = (this.tagName.toLowerCase() === "object" || (n.attr("height") && !isNaN(parseInt(n.attr("height"), 10)))) ? parseInt(n.attr("height"), 10) : n.height(),
                    k = !isNaN(parseInt(n.attr("width"), 10)) ? parseInt(n.attr("width"), 10) : n.width(),
                    l = j / k;
                if (!n.attr("id")) {
                    var m = "fitvid" + Math.floor(Math.random() * 999999);
                    n.attr("id", m)
                }
                n.wrap('<div class="fluid-width-video-wrapper"></div>').parent(".fluid-width-video-wrapper").css("padding-top", (l * 100) + "%");
                n.removeAttr("height").removeAttr("width")
            })
        })
    }
})(window.jQuery || window.Zepto);


/* __________________ 8. PLACEHOLDER SUPPORT  __________________*/
/* Placeholders.js v3.0.2 */
(function(t) {
    "use strict";

    function e(t, e, r) {
        return t.addEventListener ? t.addEventListener(e, r, !1) : t.attachEvent ? t.attachEvent("on" + e, r) : void 0
    }

    function r(t, e) {
        var r, n;
        for (r = 0, n = t.length; n > r; r++)
            if (t[r] === e) return !0;
        return !1
    }

    function n(t, e) {
        var r;
        t.createTextRange ? (r = t.createTextRange(), r.move("character", e), r.select()) : t.selectionStart && (t.focus(), t.setSelectionRange(e, e))
    }

    function a(t, e) {
        try {
            return t.type = e, !0
        } catch (r) {
            return !1
        }
    }
    t.Placeholders = {
        Utils: {
            addEventListener: e,
            inArray: r,
            moveCaret: n,
            changeType: a
        }
    }
})(this),
function(t) {
    "use strict";

    function e() {}

    function r() {
        try {
            return document.activeElement
        } catch (t) {}
    }

    function n(t, e) {
        var r, n, a = !!e && t.value !== e,
            u = t.value === t.getAttribute(V);
        return (a || u) && "true" === t.getAttribute(D) ? (t.removeAttribute(D), t.value = t.value.replace(t.getAttribute(V), ""), t.className = t.className.replace(R, ""), n = t.getAttribute(F), parseInt(n, 10) >= 0 && (t.setAttribute("maxLength", n), t.removeAttribute(F)), r = t.getAttribute(P), r && (t.type = r), !0) : !1
    }

    function a(t) {
        var e, r, n = t.getAttribute(V);
        return "" === t.value && n ? (t.setAttribute(D, "true"), t.value = n, t.className += " " + I, r = t.getAttribute(F), r || (t.setAttribute(F, t.maxLength), t.removeAttribute("maxLength")), e = t.getAttribute(P), e ? t.type = "text" : "password" === t.type && M.changeType(t, "text") && t.setAttribute(P, "password"), !0) : !1
    }

    function u(t, e) {
        var r, n, a, u, i, l, o;
        if (t && t.getAttribute(V)) e(t);
        else
            for (a = t ? t.getElementsByTagName("input") : b, u = t ? t.getElementsByTagName("textarea") : f, r = a ? a.length : 0, n = u ? u.length : 0, o = 0, l = r + n; l > o; o++) i = r > o ? a[o] : u[o - r], e(i)
    }

    function i(t) {
        u(t, n)
    }

    function l(t) {
        u(t, a)
    }

    function o(t) {
        return function() {
            m && t.value === t.getAttribute(V) && "true" === t.getAttribute(D) ? M.moveCaret(t, 0) : n(t)
        }
    }

    function c(t) {
        return function() {
            a(t)
        }
    }

    function s(t) {
        return function(e) {
            return A = t.value, "true" === t.getAttribute(D) && A === t.getAttribute(V) && M.inArray(C, e.keyCode) ? (e.preventDefault && e.preventDefault(), !1) : void 0
        }
    }

    function d(t) {
        return function() {
            n(t, A), "" === t.value && (t.blur(), M.moveCaret(t, 0))
        }
    }

    function g(t) {
        return function() {
            t === r() && t.value === t.getAttribute(V) && "true" === t.getAttribute(D) && M.moveCaret(t, 0)
        }
    }

    function v(t) {
        return function() {
            i(t)
        }
    }

    function p(t) {
        t.form && (T = t.form, "string" == typeof T && (T = document.getElementById(T)), T.getAttribute(U) || (M.addEventListener(T, "submit", v(T)), T.setAttribute(U, "true"))), M.addEventListener(t, "focus", o(t)), M.addEventListener(t, "blur", c(t)), m && (M.addEventListener(t, "keydown", s(t)), M.addEventListener(t, "keyup", d(t)), M.addEventListener(t, "click", g(t))), t.setAttribute(j, "true"), t.setAttribute(V, x), (m || t !== r()) && a(t)
    }
    var b, f, m, h, A, y, E, x, L, T, N, S, w, B = ["text", "search", "url", "tel", "email", "password", "number", "textarea"],
        C = [27, 33, 34, 35, 36, 37, 38, 39, 40, 8, 46],
        k = "#ccc",
        I = "placeholdersjs",
        R = RegExp("(?:^|\\s)" + I + "(?!\\S)"),
        V = "data-placeholder-value",
        D = "data-placeholder-active",
        P = "data-placeholder-type",
        U = "data-placeholder-submit",
        j = "data-placeholder-bound",
        q = "data-placeholder-focus",
        z = "data-placeholder-live",
        F = "data-placeholder-maxlength",
        G = document.createElement("input"),
        H = document.getElementsByTagName("head")[0],
        J = document.documentElement,
        K = t.Placeholders,
        M = K.Utils;
    if (K.nativeSupport = void 0 !== G.placeholder, !K.nativeSupport) {
        for (b = document.getElementsByTagName("input"), f = document.getElementsByTagName("textarea"), m = "false" === J.getAttribute(q), h = "false" !== J.getAttribute(z), y = document.createElement("style"), y.type = "text/css", E = document.createTextNode("." + I + " { color:" + k + "; }"), y.styleSheet ? y.styleSheet.cssText = E.nodeValue : y.appendChild(E), H.insertBefore(y, H.firstChild), w = 0, S = b.length + f.length; S > w; w++) N = b.length > w ? b[w] : f[w - b.length], x = N.attributes.placeholder, x && (x = x.nodeValue, x && M.inArray(B, N.type) && p(N));
        L = setInterval(function() {
            for (w = 0, S = b.length + f.length; S > w; w++) N = b.length > w ? b[w] : f[w - b.length], x = N.attributes.placeholder, x ? (x = x.nodeValue, x && M.inArray(B, N.type) && (N.getAttribute(j) || p(N), (x !== N.getAttribute(V) || "password" === N.type && !N.getAttribute(P)) && ("password" === N.type && !N.getAttribute(P) && M.changeType(N, "text") && N.setAttribute(P, "password"), N.value === N.getAttribute(V) && (N.value = x), N.setAttribute(V, x)))) : N.getAttribute(D) && (n(N), N.removeAttribute(V));
            h || clearInterval(L)
        }, 100)
    }
    M.addEventListener(t, "beforeunload", function() {
        K.disable()
    }), K.disable = K.nativeSupport ? e : i, K.enable = K.nativeSupport ? e : l
}(this);


/*! __________________ 9. MENU SCRIPT __________________*/
/*!* DC jQuery Vertical Accordion Menu - jQuery vertical accordion menu plugin
/*! Copyright (c) 2011 Design Chemical * Dual licensed under the MIT and GPL licenses: * http://www.opensource.org/licenses/mit-license.php * http://www.gnu.org/licenses/gpl.html */
(function(b) {
    b.fn.dcAccordion = function(h) {
        var f = {
            classParent: "parent",
            classActive: "current",
            classArrow: "dcjq-icon",
            classCount: "dcjq-count",
            classExpand: "current-parent",
            eventType: "click",
            hoverDelay: 300,
            menuClose: true,
            autoClose: true,
            autoExpand: false,
            speed: "slow",
            saveState: true,
            disableLink: true,
            showCount: false,
            cookie: "dcjq-accordion"
        };
        var h = b.extend(f, h);
        this.each(function(c) {
            var v = this;
            r();
            if (f.saveState == true) {
                g(f.cookie, v)
            }
            if (f.autoExpand == true) {
                b("li." + f.classExpand + " > a").addClass(f.classActive)
            }
            t();
            if (f.eventType == "hover") {
                var w = {
                    sensitivity: 2,
                    interval: f.hoverDelay,
                    over: d,
                    timeout: f.hoverDelay,
                    out: e
                };
                b("li a", v).hoverIntent(w);
                var x = {
                    sensitivity: 2,
                    interval: 1000,
                    over: q,
                    timeout: 1000,
                    out: u
                };
                b(v).hoverIntent(x);
                if (f.disableLink == true) {
                    b("li a", v).click(function(i) {
                        if (b(this).siblings("ul").length > 0) {
                            i.preventDefault()
                        }
                    })
                }
            } else {
                b("li a", v).click(function(i) {
                    $activeLi = b(this).parent("li");
                    $parentsLi = $activeLi.parents("li");
                    $parentsUl = $activeLi.parents("ul");
                    if (f.disableLink == true) {
                        if (b(this).siblings("ul").length > 0) {
                            i.preventDefault()
                        }
                    }
                    if (f.autoClose == true) {
                        s($parentsLi, $parentsUl)
                    }
                    if (b("> ul", $activeLi).is(":visible")) {
                        b("ul", $activeLi).slideUp(f.speed);
                        b("a", $activeLi).removeClass(f.classActive)
                    } else {
                        b(this).siblings("ul").slideToggle(f.speed);
                        b("> a", $activeLi).addClass(f.classActive)
                    }
                    if (f.saveState == true) {
                        a(f.cookie, v)
                    }
                })
            }

            function r() {
                $arrow = '<span class="' + f.classArrow + '"></span>';
                var i = f.classParent + "-li";
                b("> ul", v).show();
                b("li", v).each(function() {
                    if (b("> ul", this).length > 0) {
                        b(this).addClass(i);
                        b("> a", this).addClass(f.classParent).append($arrow)
                    }
                });
                b("> ul", v).hide();
                if (f.showCount == true) {
                    b("li." + i, v).each(function() {
                        if (f.disableLink == true) {
                            var j = parseInt(b("ul a:not(." + f.classParent + ")", this).length)
                        } else {
                            var j = parseInt(b("ul a", this).length)
                        }
                        b("> a", this).append(' <span class="' + f.classCount + '">(' + j + ")</span>")
                    })
                }
            }

            function d() {
                $activeLi = b(this).parent("li");
                $parentsLi = $activeLi.parents("li");
                $parentsUl = $activeLi.parents("ul");
                if (f.autoClose == true) {
                    s($parentsLi, $parentsUl)
                }
                if (b("> ul", $activeLi).is(":visible")) {
                    b("ul", $activeLi).slideUp(f.speed);
                    b("a", $activeLi).removeClass(f.classActive)
                } else {
                    b(this).siblings("ul").slideToggle(f.speed);
                    b("> a", $activeLi).addClass(f.classActive)
                }
                if (f.saveState == true) {
                    a(f.cookie, v)
                }
            }

            function e() {}

            function q() {}

            function u() {
                if (f.menuClose == true) {
                    b("ul", v).slideUp(f.speed);
                    b("a", v).removeClass(f.classActive);
                    a(f.cookie, v)
                }
            }

            function s(j, i) {
                b("ul", v).not(i).slideUp(f.speed);
                b("a", v).removeClass(f.classActive);
                b("> a", j).addClass(f.classActive)
            }

            function t() {
                b("ul", v).hide();
                $allActiveLi = b("a." + f.classActive, v);
                $allActiveLi.siblings("ul").show()
            }
        });

        function g(e, c) {
            var j = b.cookie(e);
            if (j != null) {
                var d = j.split(",");
                b.each(d, function(o, i) {
                    var p = b("li:eq(" + i + ")", c);
                    b("> a", p).addClass(f.classActive);
                    var n = p.parents("li");
                    b("> a", n).addClass(f.classActive)
                })
            }
        }

        function a(d, c) {
            var e = [];
            b("li a." + f.classActive, c).each(function(n) {
                var i = b(this).parent("li");
                var m = b("li", c).index(i);
                e.push(m)
            });
            b.cookie(d, e, {
                path: "/"
            })
        }
    }
})(jQuery);


/*! ------------ 10. INITIALIZE MENU------------------------- */
$(function() {

    $("#nav > ul").dcAccordion({
        saveState: false,
        autoClose: true,
        disableLink: true,
        speed: "fast",
        showCount: false,
        autoExpand: false
    });

});
// ------------------ END MENU --------------------



$(window).load(function() {

    // ------------------ 11. SOUND CLOUD NON-BLOCKING LOAD -------------------- //
    $(".soundcloud-wrapper").each(function() {
        var URL = $(this).attr('id');
        var htm = '<iframe width="100%" height="200px" src="https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/' + URL + '&amp;auto_play=false&amp;hide_related=false&amp;show_comments=false&amp;show_user=false&amp;show_reposts=false&amp;visual=true" frameborder="0"></iframe>';
        $(this).html(htm).fitVids().removeClass('.loading');
    });

    // ------------------ 12. BIG TEXT (window load gets correct size on load, doc ready does not) --------------------
    $('.big-text-banner').bigtext({
        minfontsize: 20 // default is null
    }).css('visibility', 'visible');


});


$(function() {

    /* --------------- 13. DATA SLIDE --------------- */
    $('[data-slide="slide"]').click(function(e) {

        var $this = $(this);
        var target = $this.attr('data-target');
        var $target = $(target);
        if ($('.slide-panel-parent').children().is('.open')) {
            $('.open').not(target).removeClass('open');
            $('.active-slide-btn').not(this).removeClass('active-slide-btn');
            $(this).toggleClass('active-slide-btn');
            $(target).toggleClass('open');
            $('html').removeClass('slide-active');
        } else {
            $(target).toggleClass('open');
            $(this).toggleClass('active-slide-btn');
            $('#page').toggleClass('page-off');
        }

        if ($('.slide-panel-parent').children().is('.open')) {
            $('html').addClass('slide-active'); //was addClass

        } else {
            $('html').removeClass('slide-active');

        }

        e.preventDefault();
    });

    //correct the shifting of the scrollbar when a slide is active
    if ($(document).height() > $(window).height()) {
        $('body').addClass('body-scroll-fix');
    }


    $('.slide-panel .close').click(function(e) {
        $('.active-slide-btn').removeClass('active-slide-btn');
        $(this).parent().removeClass('open');
        $('html').removeClass('slide-active');
        $('#page').removeClass('page-off');
        e.preventDefault();
    });


    // indicate what panel you're on when you've clicked inside a panel to another panel
    $('.slide-panel .signin-toggle').click(function(e) {
        $('.header-btn.signin-toggle').toggleClass('active-slide-btn');
        e.preventDefault();
    });

    $('.slide-panel .login-toggle').click(function(e) {
        $('.header-btn.login-toggle').toggleClass('active-slide-btn');
        e.preventDefault();
    });


});


// ------------------ 14. DON'T COVER ANCHOR TARGETS --------------------

$(function() {

    //set the variable of the Navbar
    var navbarHeight = $(".header").height() + 10;

    // SLIDE TO ANCHOR and DON'T COVER ANCHORS WITH .has-anchor on the trigger :: http://stackoverflow.com/a/20320919/1004312
    $('.has-anchor').click(function() {
        if (location.pathname.replace(/^\//, '') == this.pathname.replace(/^\//, '') && location.hostname == this.hostname) {
            var target = $(this.hash);
            target = target.length ? target : $('[name=' + this.hash.slice(1) + ']');
            if (target.length) {
                $('html,body').animate({
                    scrollTop: target.offset().top - navbarHeight //offset
                }, 1500);
            }
        }
    });
    // Executed on page load with URL containing an anchor tag.
    if ($(location.href.split("#")[1])) {
        var target = $('#' + location.href.split("#")[1]);
        if (target.length) {
            $('html,body').animate({
                scrollTop: target.offset().top - navbarHeight //offset height of header here if the header is fixed and it is not!
            }, 1500);
            return false;
        }
    }

});


// --------------- 15. STICKY HEADER ---------------
/* if window height greater than or equal to the height of the vertical icons and the header stick then not stick otherwise */
$(window).on("load resize", function() {

    var windowHeight = $(window).height();
    var buttonWrapper = $('.header-btn-wrapper').outerHeight();
    var headerHeight = $('.header').outerHeight();
    var goToTopHeight = $('#go-to-top').height();


    // if (windowHeight >= buttonWrapper + headerHeight) {
        // $('.header').addClass('sticky-header');
    // } else {
        // $('.header').removeClass('sticky-header');

    // }

    if (windowHeight >= buttonWrapper + headerHeight + goToTopHeight) {
        $('#go-to-top').addClass('position-1');
    } else {
        $('#go-to-top').removeClass('position-1');

    }

});


// --------------- 16. MATCH HEIGHT OF SLIDER NEXT TO THIS COLUMN ---------------

$(window).on("load resize", function() {

    var containerSize = $('.container').width();
    var colHeight = $('.featured-carousel img').innerHeight() // - 30//;

    if (containerSize >= 767) {
        $('.column-inner > .column-bg-fake').height(colHeight);
    }

});



/* --------------- 17. MISC INITIALIZATIONS AND CUSTOM SCRIPTS --------------- */
$(function() {


    /* --------------- 17. - 1 :: EQUAL HEIGHT INITIALIZE --------------- */
    $('.equal-height-content-box .equal-height-content').responsiveEqualHeightGrid();
    $('.equal-height-col [class*="col-"]').responsiveEqualHeightGrid();

    /* --------------- 17. - 2 :: .title inner wrapper --------------- */
    $('.title').wrapInner("<span></span>");

    /* --------------- 17. - 3 :: Inner Wrappers for .info-text-box and .widget-title --------------- */
    $('.info-text-box').wrapInner("<span></span>");
    $('.widget-title').wrapInner("<span></span>");

    /* --------------- 17. - 4 :: icon-click functions :: you can remove this on live sites ---------------*/
    $('.icon-demo span.glyphicon, .icon-demo span.ti').each(function() {
        var className = $(this).attr('class');
        $(this).attr('title', className).css({
            'cursor': 'pointer'
        });
        $(this).tooltip({
            trigger: 'click',
            container: 'body'
        }).toggleClass('active');
        $(this).after("<p>" + className + "</p>");
    });
    $('body').on('mouseup', function(e) {
        $('.icon-demo span.glyphicon, .icon-demo span.ti').each(function() {
            if (!$(this).is(e.target) && $(this).has(e.target).length === 0 && $('.tooltip').has(e.target).length === 0) {
                $(this).tooltip('hide');
            }
        });
    });

    /* --------------- 17. - 5 :: Scroll to Top ---------------*/
    $('#go-to-top').hide();
    $(window).scroll(function() {
        if ($(this).scrollTop() > 300) {
            $('#go-to-top').fadeIn();
        } else {
            $('#go-to-top').fadeOut();
        }
    });
    $('#go-to-top').click(function(e) {
        $('html, body').animate({
            scrollTop: 0
        }, 800);
    });


    /* --------------- 17. - 6 :: Video Non-Blocking Load ---------------*/
    $('.video-holder.youtube .video-trigger').click(function(e) {
        e.preventDefault();
        var URL = $(this).attr('href');
        var htm = '<iframe width="560" height="315" src="http://www.youtube.com/embed/' + URL + '?rel=0?wmode=transparent" frameborder="0"></iframe>';
        $(this).parent().html(htm);
        $('.video-holder').fitVids();
        return false;
    });
    $('.video-holder.vimeo .video-trigger').click(function(e) {
        e.preventDefault();
        var URL = $(this).attr('href');
        var htm = '<iframe width="500" height="281" src="http://player.vimeo.com/video/' + URL + '?title=0&amp;byline=0&amp;portrait=0?wmode=transparent" frameborder="0"></iframe>';
        $(this).parent().html(htm);
        $('.video-holder').fitVids();
        return false;
    });

    //STOP VIDEO inside slider on click inside the slider
    $('.content-slider, .blog-slider, .featured-slider').click(function() {
        var element = $(this).find('.video-holder iframe');
        if (element.is(':visible')) {
            $.fn.videoStopFunction();
        }
    });

    $('.content-slider, .blog-slider, .featured-slider').on('touchstart', function() {
        var element = $(this).find('.video-holder iframe');
        if (element.is(':visible')) {
            $.fn.videoStopFunction();
        }
    });

    $.fn.videoStopFunction = function() {
        var video = $(".video-holder iframe").attr("src");
        $(".video-holder iframe").attr("src", "");
        $(".video-holder iframe").attr("src", video);
    }


    /* --------------- 17. - 7 :: pricing plan ---------------*/
    $('.the-details').hide();
    $('.has-details').click(function() {
        if ($(this).hasClass('active')) {
            $(this).removeClass('active');
            $(this).children('.the-details').slideUp();
        } else {
            $('.has-details').removeClass('active');
            $(this).addClass('active');
            $('.the-details').slideUp();
            $(this).children('.the-details').slideDown();
        }
        return false;
    });
    $(document).bind("mouseup touchend", function(a) {
        if ($(a.target).parents().index($(".price-row")) == -1) {
            $('.has-details').removeClass('active');
            $('.the-details').slideUp();
        }
    });

    /* --------------- 17. - 8 :: CLICK SLIDE (mouseover on no-touch) ---------------*/
    var delay = (function() {
        var timer = 0;
        return function(callback, ms) {
            clearTimeout(timer);
            timer = setTimeout(callback, ms);
        };
    })();
    $('.click-slide').show();
    if ($('html').hasClass('touch')) {
        $('.click-slide').click(function() {
            if ($(this).hasClass('open')) {
                $(this).removeClass('open');
            } else {
                $('.click-slide').removeClass('open');
                $(this).addClass('open');
            }
        });
    } else {
        $('.click-slide').hover(function() { // Hover Intent is being used here
            if ($(this).hasClass('open')) {
                $(this).removeClass('open');
            } else {
                $('.click-slide').removeClass('open');
                $(this).addClass('open');
            }
        });
    }

    $(document).bind("mouseup touchend", function(e) {
        if ($(e.target).parents(".click-slide").length === 0) {
            $('.click-slide').removeClass('open');
        }
    });

}); // end doc ready



/* --------------- 18. - STICKY ADD THIS AND SIDEBAR FOLLOW ---------------*/
// very modified -- http://www.keanrichmond.com/a-simple-sticky-sidebar.html
     if ($(".addthis_inner").length > 0) {
         var $entry = $(".entry-post"),
             $addthis = $(".addthis_inner"),
             $window2 = $(window),
             $document2 = $(document);

         function stickyAddThis() {
             var position = $('.primary-column').position().left + 15;
             var headerHeight = $('.header').outerHeight() + 30;
             var colwidth = $('.addthis_inner').outerWidth(true);
             $('.addthis_inner').css('width', (colwidth) + 'px');
             $('.addthis_inner').css('top', (headerHeight) + 'px');
             $addthis.css('left', (position) + 'px');

             var fixedBottom = $document2.height() - ($entry.offset().top + $entry.height());
             var topOffset = $entry.offset().top;
             var stopPoint = $document2.height() - (fixedBottom + $addthis.height() + 40);
             if ($entry.height() > $addthis.height() + 0) {
                 if ($window.scrollTop() > topOffset) {
                     $addthis.addClass("sticky-addthis");
                     if ($window.scrollTop() > stopPoint) {
                         $addthis.css("top", (stopPoint - $window2.scrollTop() + 40));
                     }
                 } else {
                     $addthis.removeClass("sticky-addthis");
                     $addthis.removeAttr('style');
                 }
             }
         }

     }
 if ($(".blog-single-sidebar").length > 0) {
     var $content = $(".primary-column"),
         $sidebar = $(".blog-single-sidebar"),
         $window = $(window),
         $document = $(document);

     function stickySidebar() {
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
                 $sidebar.addClass("sticky-sidebar");
                 if ($window.scrollTop() > stopPoint) {
                     $sidebar.css("top", (stopPoint - $window.scrollTop() + 40));
                 }
             } else {
                 $sidebar.removeClass("sticky-sidebar");
                 $sidebar.removeAttr('style');
             }
         }
     }
     stickySidebar();
     stickyAddThis();
     window.onscroll = function() {
         stickySidebar();
         stickyAddThis();
     };
     window.onresize = function() {
         stickySidebar();
         stickyAddThis();
     };
 }


/* --------------- 19. - SUMMARY PAGE SIDEBAR FOLLOW ---------------*/
    if ($(".blog-summary-sidebar").length > 0) {
        var $content3 = $(".primary-column"),
            $sidebar3 = $(".blog-summary-sidebar"),
            $window3 = $(window),
            $document3 = $(document);

        function stickySidebar2() {
            var position = $('#sidebar').position().left + 15;
            var headerHeight = $('.header').outerHeight();
            var colwidth = $sidebar3.outerWidth(true);
            $sidebar3.css('width', (colwidth) + 'px');
            $sidebar3.css('top', (headerHeight) + 'px');
            $sidebar3.css('left', (position) + 'px');

            var fixedBottom = $document3.height() - ($content3.offset().top + $content3.height());
            var topOffset = $content3.offset().top;
            var stopPoint = $document3.height() - (fixedBottom + $sidebar3.height() + 40);
            if ($content3.height() > $sidebar3.height() + 0) {
                if ($window3.scrollTop() > topOffset) {
                    $sidebar3.addClass("sticky-sidebar");
                    if ($window3.scrollTop() > stopPoint) {
                        $sidebar3.css("top", (stopPoint - $window3.scrollTop() + 40));
                    }
                } else {
                    $sidebar3.removeClass("sticky-sidebar");
                    $sidebar3.removeAttr('style');
                }
            }
        }
        stickySidebar2();
        window.onscroll = function() {
            stickySidebar2();
        };
        window.onresize = function() {
            stickySidebar2();
        };
    }

/* --------------- 20. - JQUERY STICKY FOOTER ---------------*/
/*!
 * jQuery Sticky Footer v1.2.3
 * Copyright 2014 miWebb
 * https://github.com/miWebb/jQuery.stickyFooter
 * Released under the MIT license
 */
(function(e, t, n) {
    "use strict";

    function r(t, n) {
        var r = e(t).outerHeight(true);
        e(n.content).each(function() {
            r += e(this).outerHeight(true)
        });
        if (r < e(n.frame ? n.frame : t.parent()).height()) {
            e(t).addClass(n.class)
        } else {
            e(t).removeClass(n.class)
        }
    }
    e.fn.stickyFooter = function(n) {
        var n = e.extend({}, e.fn.stickyFooter.defaults, n);
        var i = this;
        r(i, n);
        e(t).resize(function() {
            r(i, n)
        });
        return this
    };
    e.fn.stickyFooter.defaults = {
        "class": "sticky-footer",
        frame: "",
        content: "#page"
    }
})(jQuery, window);

$(window).on('load resize', function() {
   $(".footer").stickyFooter();
});
