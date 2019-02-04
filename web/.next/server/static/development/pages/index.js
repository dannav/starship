module.exports =
/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = require('../../../ssr-module-cache.js');
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		var threw = true;
/******/ 		try {
/******/ 			modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/ 			threw = false;
/******/ 		} finally {
/******/ 			if(threw) delete installedModules[moduleId];
/******/ 		}
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, { enumerable: true, get: getter });
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 			Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 		}
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// create a fake namespace object
/******/ 	// mode & 1: value is a module id, require it
/******/ 	// mode & 2: merge all properties of value into the ns
/******/ 	// mode & 4: return value when already ns object
/******/ 	// mode & 8|1: behave like require
/******/ 	__webpack_require__.t = function(value, mode) {
/******/ 		if(mode & 1) value = __webpack_require__(value);
/******/ 		if(mode & 8) return value;
/******/ 		if((mode & 4) && typeof value === 'object' && value && value.__esModule) return value;
/******/ 		var ns = Object.create(null);
/******/ 		__webpack_require__.r(ns);
/******/ 		Object.defineProperty(ns, 'default', { enumerable: true, value: value });
/******/ 		if(mode & 2 && typeof value != 'string') for(var key in value) __webpack_require__.d(ns, key, function(key) { return value[key]; }.bind(null, key));
/******/ 		return ns;
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = 3);
/******/ })
/************************************************************************/
/******/ ({

/***/ "./components/homepage.js":
/*!********************************!*\
  !*** ./components/homepage.js ***!
  \********************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var styled_jsx_style__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! styled-jsx/style */ "styled-jsx/style");
/* harmony import */ var styled_jsx_style__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(styled_jsx_style__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ "react");
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var react_redux__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! react-redux */ "react-redux");
/* harmony import */ var react_redux__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(react_redux__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! next/link */ "next/link");
/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(next_link__WEBPACK_IMPORTED_MODULE_3__);


function _typeof(obj) { if (typeof Symbol === "function" && typeof Symbol.iterator === "symbol") { _typeof = function _typeof(obj) { return typeof obj; }; } else { _typeof = function _typeof(obj) { return obj && typeof Symbol === "function" && obj.constructor === Symbol && obj !== Symbol.prototype ? "symbol" : typeof obj; }; } return _typeof(obj); }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; var ownKeys = Object.keys(source); if (typeof Object.getOwnPropertySymbols === 'function') { ownKeys = ownKeys.concat(Object.getOwnPropertySymbols(source).filter(function (sym) { return Object.getOwnPropertyDescriptor(source, sym).enumerable; })); } ownKeys.forEach(function (key) { _defineProperty(target, key, source[key]); }); } return target; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); return Constructor; }

function _possibleConstructorReturn(self, call) { if (call && (_typeof(call) === "object" || typeof call === "function")) { return call; } return _assertThisInitialized(self); }

function _getPrototypeOf(o) { _getPrototypeOf = Object.setPrototypeOf ? Object.getPrototypeOf : function _getPrototypeOf(o) { return o.__proto__ || Object.getPrototypeOf(o); }; return _getPrototypeOf(o); }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function"); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, writable: true, configurable: true } }); if (superClass) _setPrototypeOf(subClass, superClass); }

function _setPrototypeOf(o, p) { _setPrototypeOf = Object.setPrototypeOf || function _setPrototypeOf(o, p) { o.__proto__ = p; return o; }; return _setPrototypeOf(o, p); }

function _assertThisInitialized(self) { if (self === void 0) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return self; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }





var Home =
/*#__PURE__*/
function (_React$Component) {
  _inherits(Home, _React$Component);

  function Home(props) {
    var _this;

    _classCallCheck(this, Home);

    _this = _possibleConstructorReturn(this, _getPrototypeOf(Home).call(this, props));

    _defineProperty(_assertThisInitialized(_assertThisInitialized(_this)), "state", {
      termTab: 'search'
    });

    _this.setTermTab = _this.setTermTab.bind(_assertThisInitialized(_assertThisInitialized(_this)));
    return _this;
  }

  _createClass(Home, [{
    key: "setTermTab",
    value: function setTermTab(tab) {
      this.setState(function (old) {
        return _objectSpread({}, old, {
          termTab: tab
        });
      });
    }
  }, {
    key: "render",
    value: function render() {
      var _this2 = this;

      return react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(react__WEBPACK_IMPORTED_MODULE_1___default.a.Fragment, null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "full-width"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "stars"
      }), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "twinkling"
      }), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "clouds"
      }), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "hero"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "container"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("h2", null, "Keep Your Remote Team In Sync"), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("p", null, "Starship is a private search engine for you and your team. Make finding the information anyone on your team needs one command away."), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "hero-actions"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("button", {
        type: "button",
        className: "btn green"
      }, "Starship in 60 seconds")), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("button", {
        type: "button",
        className: "btn red"
      }, "Get started - It's free"))))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "container tab-section pull-front"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("ul", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("li", {
        className: this.state.termTab === 'search' ? 'active' : null
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("button", {
        type: "button",
        onClick: this.setTermTab.bind(this, 'search')
      }, "Search")), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("li", {
        className: this.state.termTab === 'index' ? 'active' : null
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("button", {
        type: "button",
        onClick: this.setTermTab.bind(this, 'index')
      }, "Index"))))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "showcase full-width"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "container"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "terminal",
        id: "termynal",
        "data-terminal": true
      }, function () {
        var url = '/static/search.svg' + "?" + Math.random();

        if (_this2.state.termTab === 'search') {
          return react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("img", {
            src: url
          });
        }

        url = '/static/index.svg' + "?" + Math.random();
        return react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("img", {
          src: url
        });
      }()))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "showcase full-width alt"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "container"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("h2", null, "Your Private Search Engine"), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("p", null, "Sifting through project documentation, company information, and your personal notes. Starship lets your team upload any document and search them with one command."), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "two-sided"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "placeholder"
      }), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "description"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("p", null, "Starship was built to make sharing documented information in remote teams easier."), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("p", null, "Store PDF's, Word documents, HTML pages, Markdown, or just about any other text based file format."), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("p", null, "We'll index the content and make it searchable so that you can share it with others and find all of your team's information later."))))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "showcase full-width"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "container three"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("h2", null, "AI-Powered"), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("p", {
        className: "hug"
      }, "Starship uses AI to provide better results than your typical search.")), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("h2", null, "Security Focused"), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("p", {
        className: "hug"
      }, "All documents stored are private and only accessible by your team.")), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("h2", null, "Command Line First"), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("p", {
        className: "hug"
      }, "Index, search, and download documents from a CLI.")))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "showcase full-width alt"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "container cta-footer"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("h2", null, "Try Starship For Free")), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
        className: "actions"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("a", {
        className: "btn red"
      }, "Get Started"), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("a", {
        className: "btn green"
      }, "Starship in 60 seconds")))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(styled_jsx_style__WEBPACK_IMPORTED_MODULE_0___default.a, {
        styleId: "1296775629",
        css: ".pull-front,.pull-front *{z-index:999;}.stars,.twinkling,.clouds{position:absolute;display:block;top:0;bottom:0;left:0;right:0;width:100%;height:100%;}.stars{z-index:0;background:#FFF url('/static/stars.png') repeat top center;}.twinkling{z-index:1;background:transparent url('/static/twinkling.png') repeat top center;-webkit-animation:move-twink-back 200s linear infinite;animation:move-twink-back 200s linear infinite;}.clouds{z-index:2;background:transparent url('/static/clouds.png') repeat top center;-webkit-animation:move-clouds-back 200s linear infinite;animation:move-clouds-back 200s linear infinite;}@-webkit-keyframes move-twink-back{from{background-position:0 0;}to{background-position:-10000px 5000px;}}@keyframes move-twink-back{from{background-position:0 0;}to{background-position:-10000px 5000px;}}@-webkit-keyframes move-clouds-back{from{background-position:0 0;}to{background-position:10000px 0;}}@keyframes move-clouds-back{from{background-position:0 0;}to{background-position:10000px 0;}}[data-terminal]{max-width:100%;background:var(--color-bg);border-radius:4px;padding:40px 15px 20px 15px;position:relative;-webkit-box-sizing:border-box;box-sizing:border-box;}[data-terminal]:before{content:'';position:absolute;top:15px;left:15px;display:inline-block;width:15px;height:15px;border-radius:50%;background:#d9515d;-webkit-box-shadow:25px 0 0 #f4c025,50px 0 0 #3ec930;box-shadow:25px 0 0 #f4c025,50px 0 0 #3ec930;}.cta-footer{display:grid;grid-template-columns:.75fr 1.25fr;justify-items:center;-webkit-align-items:center;-webkit-box-align:center;-ms-flex-align:center;align-items:center;}.cta-footer h2{text-align:left !important;margin:0;}.cta-footer .btn:first-child{margin-right:20px;}.cta-footer .btn{box-shadow:0 20px 50px 0 rgba(0,0,0,0.2);}.hero .btn{box-shadow:0 20px 50px 0 rgba(0,0,0,0.2);}.three{display:grid;grid-template-columns:1fr 1fr 1fr;-webkit-column-gap:1.5rem;column-gap:1.5rem;}.three div{justify-self:center;}.tab-section{display:grid;grid-template-columns:1fr;}.three h2{margin:0;text-align:left !important;font-size:100% !important;}.tab-section ul{justify-self:center;list-style:none;margin:0;padding:0;}.tab-section li{display:inline-block;float:left;padding:5px 20px;}.tab-section li.active{border-bottom:1px solid #122239;}.tacb-section li:hover{cursor:pointer;}.tab-section button{outline:none;border:none;color:#122239;}.tab-section button:hover{color:#122239;cursor:pointer;}.tab-section button{font-size:90% !important;}.two-sided{margin:3rem 0;display:grid;grid-template-columns:1fr 1fr;-webkit-column-gap:1rem;column-gap:1rem;}.two-sided .placeholder{height:100%;width:100%;background:#122239;}.two-sided p{margin:1rem 0;}.hero{z-index:999;padding:100px 0 150px 0;width:90%;position:relative;left:5%;font-size:132%;text-align:center;}.hero h2{font-size:2.121rem;margin-bottom:1.5rem;}.hero-actions{font-size:84%;margin:40px 0 0 0;display:grid;-webkit-column-gap:1rem;column-gap:1rem;grid-template-columns:1fr 1fr;-webkit-align-items:center;-webkit-box-align:center;-ms-flex-align:center;align-items:center;}.hero-actions div:first-child{justify-self:end;}.hero-actions div:last-child{justify-self:start;}.showcase{background:#fafbfc;padding:3rem 0;border-bottom:1px solid #eaeaea;}.showcase h2{text-align:center;}.showcase.alt{background:#fff;}.showcase.alt p{color:#707d8e;}.full-width{width:100vw;position:relative;left:50%;right:50%;margin-left:-50vw;margin-right:-50vw;}.long-width{position:relative;width:calc(100% + 146px);margin-left:-73px;margin-right:-73px;}.terminal{margin:0 auto;background:#122239;box-shadow:0 20px 50px 0 rgba(0,0,0,0.2);border-radius:5px;}.terminal img{width:100%;}\n/*# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIi9Vc2Vycy9kZXYvZ28vc3JjL2dpdGh1Yi5jb20vZGFubmF2L3N0YXJzaGlwL3dlYi9jb21wb25lbnRzL2hvbWVwYWdlLmpzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiJBQWdJMkIsQUFHeUIsQUFJSyxBQVFQLEFBS0EsQUFNQSxBQU1rQixBQUNVLEFBSVYsQUFDSSxBQUlmLEFBVUosQUFlQSxBQU9jLEFBS1QsQUFJdUIsQUFJQSxBQUk1QixBQU1PLEFBSVAsQUFLSixBQU1XLEFBT0MsQUFNVyxBQUlqQixBQUlGLEFBTUMsQUFLVyxBQUlYLEFBT0YsQUFNRSxBQUlGLEFBVU8sQUFLTCxBQVNHLEFBSUUsQUFJQSxBQU1ELEFBSUYsQUFJRixBQUlGLEFBU00sQUFPSixBQU9ILFNBbklnQixDQTNGZ0MsQUFLVSxBQU1ILENBMEI5QyxBQTBMdEIsQ0EzT0EsQUF3SmEsQUFVYSxBQWtETixDQXBKaUIsQUF3QkQsQUFVUixBQWdDZCxDQU1HLEFBU0YsQUFhZixBQW1Cb0IsQUErQnBCLEFBb0JxQixDQTdMVSxBQXVGL0IsQ0E4RUEsQ0FsQkEsQ0F0TGUsQUF3RWYsQUE0SEEsQUFxQjJCLENBakRKLEFBa0J2QixBQUlpQixDQXBHakIsQUFla0IsQ0FPTCxFQXFDUSxDQTdIVSxBQUtBLENBa0dmLEFBVWhCLEVBdEVXLEFBMkVxQixFQWhHbkIsQUF1RmIsQ0F2R21DLEFBbUx4QixFQWhOSixBQWdIWSxBQUtuQixBQXdEZSxDQW1ENEIsQ0FsQ1QsRUF0S08sQUEyQ3pDLEFBZ0M0QixBQU1qQixBQXFEQyxFQTlKSSxBQThDRixDQWdEZCxBQWlDQSxBQWtGWSxDQXpDWixDQTdGQSxBQUlBLENBNUNzQixBQWtIdEIsQ0FvRW9CLEVBL0dSLEFBb0VNLENBZkUsQ0E5SlosQUFtRlksQ0FyQ0ssQUFhRixDQXFEdkIsQUFpR29CLEtBak5KLENBMEdoQixFQW1Da0IsR0ExR2MsQ0FzTFgsQ0F4TlQsQUFrR1osRUE0RFUsRUErQlYsQ0FvQnFCLEVBM01yQixBQXdDZSxBQWFNLEdBb0dKLENBL0pPLENBK05KLEdBOU04QixHQU5ELEFBb0NqQyxBQTBLaEIsS0F4TkEsQUE0S2dDLENBcUNoQyxDQWxEb0IsQ0E3SEUsR0FnRHRCLENBbkNzQixBQWdMdEIsS0FuRkEsUUFvQkEsQ0E3SGtDLElBY1gsS0E0SEYsY0EzSHFDLE9BZHhCLHNCQUNsQyxJQXNCQSxtQkExQ0EsQ0FOQSxBQXdDMEQsMEJBMkgxRCxtQkExSEEiLCJmaWxlIjoiL1VzZXJzL2Rldi9nby9zcmMvZ2l0aHViLmNvbS9kYW5uYXYvc3RhcnNoaXAvd2ViL2NvbXBvbmVudHMvaG9tZXBhZ2UuanMiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgUmVhY3QgZnJvbSAncmVhY3QnXG5pbXBvcnQgeyBjb25uZWN0IH0gZnJvbSAncmVhY3QtcmVkdXgnXG5pbXBvcnQgTGluayBmcm9tICduZXh0L2xpbmsnXG5cbmNsYXNzIEhvbWUgZXh0ZW5kcyBSZWFjdC5Db21wb25lbnQge1xuICBzdGF0ZSA9IHt0ZXJtVGFiOiAnc2VhcmNoJ31cblxuICBjb25zdHJ1Y3Rvcihwcm9wcykge1xuICAgIHN1cGVyKHByb3BzKVxuXG4gICAgdGhpcy5zZXRUZXJtVGFiID0gdGhpcy5zZXRUZXJtVGFiLmJpbmQodGhpcylcbiAgfVxuXG4gIHNldFRlcm1UYWIodGFiKSB7XG4gICAgdGhpcy5zZXRTdGF0ZShvbGQgPT4ge1xuICAgICAgcmV0dXJuIHtcbiAgICAgICAgLi4ub2xkLFxuICAgICAgICB0ZXJtVGFiOiB0YWJcbiAgICAgIH1cbiAgICB9KVxuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiAoXG4gICAgICA8PlxuICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cImZ1bGwtd2lkdGhcIj5cbiAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cInN0YXJzXCI+PC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJ0d2lua2xpbmdcIj48L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cImNsb3Vkc1wiPjwvZGl2PlxuXG4gICAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJoZXJvXCI+XG4gICAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cImNvbnRhaW5lclwiPlxuICAgICAgICAgICAgICA8aDI+XG4gICAgICAgICAgICAgICAgS2VlcCBZb3VyIFJlbW90ZSBUZWFtIEluIFN5bmNcbiAgICAgICAgICAgICAgPC9oMj5cbiAgICAgICAgICAgICAgPHA+XG4gICAgICAgICAgICAgICAgey8qIEFuIGVmZmVjdGl2ZSB0ZWFtIHJlcXVpcmVzIGNvbnN0YW50IGFjY2VzcyB0byBpbnRlcm5hbCBpbmZvcm1hdGlvbi4gUHJvdmlkZSB0aGVtIHdpdGggdGhlIHRvb2xzIHRvIGZpbmQgd2hhdGV2ZXIgdGhleSBuZWVkLCB3aGVuZXZlciB0aGV5IG5lZWQgaXQuICovfVxuICAgICAgICAgICAgICAgIFN0YXJzaGlwIGlzIGEgcHJpdmF0ZSBzZWFyY2ggZW5naW5lIGZvciB5b3UgYW5kIHlvdXIgdGVhbS4gTWFrZSBmaW5kaW5nIHRoZSBpbmZvcm1hdGlvbiBhbnlvbmUgb24geW91ciB0ZWFtIG5lZWRzIG9uZSBjb21tYW5kIGF3YXkuXG4gICAgICAgICAgICAgIDwvcD5cbiAgICAgICAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJoZXJvLWFjdGlvbnNcIj5cbiAgICAgICAgICAgICAgICA8ZGl2PlxuICAgICAgICAgICAgICAgICAgPGJ1dHRvbiB0eXBlPVwiYnV0dG9uXCIgY2xhc3NOYW1lPVwiYnRuIGdyZWVuXCI+U3RhcnNoaXAgaW4gNjAgc2Vjb25kczwvYnV0dG9uPlxuICAgICAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgICAgIDxkaXY+XG4gICAgICAgICAgICAgICAgICA8YnV0dG9uIHR5cGU9XCJidXR0b25cIiBjbGFzc05hbWU9XCJidG4gcmVkXCI+R2V0IHN0YXJ0ZWQgLSBJdCdzIGZyZWU8L2J1dHRvbj5cbiAgICAgICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cImNvbnRhaW5lciB0YWItc2VjdGlvbiBwdWxsLWZyb250XCI+XG4gICAgICAgICAgICA8dWw+XG4gICAgICAgICAgICAgIDxsaSBjbGFzc05hbWU9e3RoaXMuc3RhdGUudGVybVRhYiA9PT0gJ3NlYXJjaCcgPyAnYWN0aXZlJyA6IG51bGx9PlxuICAgICAgICAgICAgICAgIDxidXR0b24gdHlwZT1cImJ1dHRvblwiIG9uQ2xpY2s9e3RoaXMuc2V0VGVybVRhYi5iaW5kKHRoaXMsICdzZWFyY2gnKX0+U2VhcmNoPC9idXR0b24+XG4gICAgICAgICAgICAgIDwvbGk+XG4gICAgICAgICAgICAgIDxsaSBjbGFzc05hbWU9e3RoaXMuc3RhdGUudGVybVRhYiA9PT0gJ2luZGV4JyA/ICdhY3RpdmUnIDogbnVsbH0+XG4gICAgICAgICAgICAgICAgPGJ1dHRvbiB0eXBlPVwiYnV0dG9uXCIgb25DbGljaz17dGhpcy5zZXRUZXJtVGFiLmJpbmQodGhpcywgJ2luZGV4Jyl9PkluZGV4PC9idXR0b24+XG4gICAgICAgICAgICAgIDwvbGk+XG4gICAgICAgICAgICA8L3VsPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJzaG93Y2FzZSBmdWxsLXdpZHRoXCI+XG4gICAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJjb250YWluZXJcIj5cbiAgICAgICAgICAgIDxkaXYgY2xhc3NOYW1lPVwidGVybWluYWxcIiBpZD1cInRlcm15bmFsXCIgZGF0YS10ZXJtaW5hbD5cbiAgICAgICAgICAgICAge1xuICAgICAgICAgICAgICAgICgoKSA9PiB7XG4gICAgICAgICAgICAgICAgICBsZXQgdXJsID0gJy9zdGF0aWMvc2VhcmNoLnN2ZycgKyBcIj9cIiArIE1hdGgucmFuZG9tKClcbiAgICAgICAgICAgICAgICAgIGlmICh0aGlzLnN0YXRlLnRlcm1UYWIgPT09ICdzZWFyY2gnKSB7XG4gICAgICAgICAgICAgICAgICAgIHJldHVybiA8aW1nIHNyYz17dXJsfSAvPlxuICAgICAgICAgICAgICAgICAgfVxuXG4gICAgICAgICAgICAgICAgICB1cmwgPSAnL3N0YXRpYy9pbmRleC5zdmcnICsgXCI/XCIgKyBNYXRoLnJhbmRvbSgpXG4gICAgICAgICAgICAgICAgICByZXR1cm4gPGltZyBzcmM9e3VybH0gLz5cbiAgICAgICAgICAgICAgICB9KSgpXG4gICAgICAgICAgICAgIH1cbiAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJzaG93Y2FzZSBmdWxsLXdpZHRoIGFsdFwiPlxuICAgICAgICAgIDxkaXYgY2xhc3NOYW1lPVwiY29udGFpbmVyXCI+XG4gICAgICAgICAgICA8aDI+WW91ciBQcml2YXRlIFNlYXJjaCBFbmdpbmU8L2gyPlxuICAgICAgICAgICAgPHA+XG4gICAgICAgICAgICAgIFNpZnRpbmcgdGhyb3VnaCBwcm9qZWN0IGRvY3VtZW50YXRpb24sIGNvbXBhbnkgaW5mb3JtYXRpb24sIGFuZCB5b3VyIHBlcnNvbmFsIG5vdGVzLiBTdGFyc2hpcCBsZXRzIHlvdXIgdGVhbSB1cGxvYWQgYW55IGRvY3VtZW50IGFuZCBzZWFyY2ggdGhlbSB3aXRoIG9uZSBjb21tYW5kLlxuICAgICAgICAgICAgPC9wPlxuICAgICAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJ0d28tc2lkZWRcIj5cbiAgICAgICAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJwbGFjZWhvbGRlclwiPjwvZGl2PlxuICAgICAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cImRlc2NyaXB0aW9uXCI+XG4gICAgICAgICAgICAgICAgPHA+XG4gICAgICAgICAgICAgICAgICBTdGFyc2hpcCB3YXMgYnVpbHQgdG8gbWFrZSBzaGFyaW5nIGRvY3VtZW50ZWQgaW5mb3JtYXRpb24gaW4gcmVtb3RlIHRlYW1zIGVhc2llci5cbiAgICAgICAgICAgICAgICA8L3A+XG4gICAgICAgICAgICAgICAgPHA+XG4gICAgICAgICAgICAgICAgICBTdG9yZSBQREYncywgV29yZCBkb2N1bWVudHMsIEhUTUwgcGFnZXMsIE1hcmtkb3duLCBvciBqdXN0IGFib3V0IGFueSBvdGhlciB0ZXh0IGJhc2VkIGZpbGUgZm9ybWF0LlxuICAgICAgICAgICAgICAgIDwvcD5cbiAgICAgICAgICAgICAgICA8cD5cbiAgICAgICAgICAgICAgICAgIFdlJ2xsIGluZGV4IHRoZSBjb250ZW50IGFuZCBtYWtlIGl0IHNlYXJjaGFibGUgc28gdGhhdCB5b3UgY2FuIHNoYXJlIGl0IHdpdGggb3RoZXJzIGFuZCBmaW5kIGFsbCBvZiB5b3VyIHRlYW0ncyBpbmZvcm1hdGlvbiBsYXRlci5cbiAgICAgICAgICAgICAgICA8L3A+XG4gICAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cInNob3djYXNlIGZ1bGwtd2lkdGhcIj5cbiAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cImNvbnRhaW5lciB0aHJlZVwiPlxuICAgICAgICAgICAgPGRpdj5cbiAgICAgICAgICAgICAgPGgyPkFJLVBvd2VyZWQ8L2gyPlxuICAgICAgICAgICAgICA8cCBjbGFzc05hbWU9XCJodWdcIj5cbiAgICAgICAgICAgICAgICBTdGFyc2hpcCB1c2VzIEFJIHRvIHByb3ZpZGUgYmV0dGVyIHJlc3VsdHMgdGhhbiB5b3VyIHR5cGljYWwgc2VhcmNoLlxuICAgICAgICAgICAgICA8L3A+XG4gICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICAgIDxkaXY+XG4gICAgICAgICAgICAgIDxoMj5TZWN1cml0eSBGb2N1c2VkPC9oMj5cbiAgICAgICAgICAgICAgPHAgY2xhc3NOYW1lPVwiaHVnXCI+QWxsIGRvY3VtZW50cyBzdG9yZWQgYXJlIHByaXZhdGUgYW5kIG9ubHkgYWNjZXNzaWJsZSBieSB5b3VyIHRlYW0uPC9wPlxuICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICA8ZGl2PlxuICAgICAgICAgICAgICA8aDI+Q29tbWFuZCBMaW5lIEZpcnN0PC9oMj5cbiAgICAgICAgICAgICAgPHAgY2xhc3NOYW1lPVwiaHVnXCI+SW5kZXgsIHNlYXJjaCwgYW5kIGRvd25sb2FkIGRvY3VtZW50cyBmcm9tIGEgQ0xJLjwvcD5cbiAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJzaG93Y2FzZSBmdWxsLXdpZHRoIGFsdFwiPlxuICAgICAgICAgIDxkaXYgY2xhc3NOYW1lPVwiY29udGFpbmVyIGN0YS1mb290ZXJcIj5cbiAgICAgICAgICAgIDxkaXY+XG4gICAgICAgICAgICAgIDxoMj5UcnkgU3RhcnNoaXAgRm9yIEZyZWU8L2gyPlxuICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cImFjdGlvbnNcIj5cbiAgICAgICAgICAgICAgPGEgY2xhc3NOYW1lPVwiYnRuIHJlZFwiPkdldCBTdGFydGVkPC9hPlxuICAgICAgICAgICAgICA8YSBjbGFzc05hbWU9XCJidG4gZ3JlZW5cIj5TdGFyc2hpcCBpbiA2MCBzZWNvbmRzPC9hPlxuICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8c3R5bGUganN4IGdsb2JhbD57YFxuICAgICAgICAgIC5wdWxsLWZyb250LCAucHVsbC1mcm9udCAqIHtcbiAgICAgICAgICAgIHotaW5kZXg6IDk5OTtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAuc3RhcnMsIC50d2lua2xpbmcsIC5jbG91ZHMge1xuICAgICAgICAgICAgcG9zaXRpb246YWJzb2x1dGU7XG4gICAgICAgICAgICBkaXNwbGF5OmJsb2NrO1xuICAgICAgICAgICAgdG9wOjA7IGJvdHRvbTowO1xuICAgICAgICAgICAgbGVmdDowOyByaWdodDowO1xuICAgICAgICAgICAgd2lkdGg6MTAwJTsgaGVpZ2h0OjEwMCU7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLnN0YXJzIHtcbiAgICAgICAgICAgIHotaW5kZXg6IDA7XG4gICAgICAgICAgICBiYWNrZ3JvdW5kOiAjRkZGIHVybCgnL3N0YXRpYy9zdGFycy5wbmcnKSByZXBlYXQgdG9wIGNlbnRlcjtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAudHdpbmtsaW5ne1xuICAgICAgICAgICAgei1pbmRleDogMTtcbiAgICAgICAgICAgIGJhY2tncm91bmQ6dHJhbnNwYXJlbnQgdXJsKCcvc3RhdGljL3R3aW5rbGluZy5wbmcnKSByZXBlYXQgdG9wIGNlbnRlcjtcbiAgICAgICAgICAgIGFuaW1hdGlvbjogbW92ZS10d2luay1iYWNrIDIwMHMgbGluZWFyIGluZmluaXRlO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC5jbG91ZHN7XG4gICAgICAgICAgICB6LWluZGV4OiAyO1xuICAgICAgICAgICAgYmFja2dyb3VuZDp0cmFuc3BhcmVudCB1cmwoJy9zdGF0aWMvY2xvdWRzLnBuZycpIHJlcGVhdCB0b3AgY2VudGVyO1xuICAgICAgICAgICAgYW5pbWF0aW9uOiBtb3ZlLWNsb3Vkcy1iYWNrIDIwMHMgbGluZWFyIGluZmluaXRlO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIEBrZXlmcmFtZXMgbW92ZS10d2luay1iYWNrIHtcbiAgICAgICAgICAgIGZyb20ge2JhY2tncm91bmQtcG9zaXRpb246MCAwO31cbiAgICAgICAgICAgIHRvIHtiYWNrZ3JvdW5kLXBvc2l0aW9uOi0xMDAwMHB4IDUwMDBweDt9XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgQGtleWZyYW1lcyBtb3ZlLWNsb3Vkcy1iYWNrIHtcbiAgICAgICAgICAgIGZyb20ge2JhY2tncm91bmQtcG9zaXRpb246MCAwO31cbiAgICAgICAgICAgIHRvIHtiYWNrZ3JvdW5kLXBvc2l0aW9uOjEwMDAwcHggMDt9XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgW2RhdGEtdGVybWluYWxdIHtcbiAgICAgICAgICAgICAgbWF4LXdpZHRoOiAxMDAlO1xuICAgICAgICAgICAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jb2xvci1iZyk7XG4gICAgICAgICAgICAgIGJvcmRlci1yYWRpdXM6IDRweDtcbiAgICAgICAgICAgICAgcGFkZGluZzogNDBweCAxNXB4IDIwcHggMTVweDtcbiAgICAgICAgICAgICAgcG9zaXRpb246IHJlbGF0aXZlO1xuICAgICAgICAgICAgICAtd2Via2l0LWJveC1zaXppbmc6IGJvcmRlci1ib3g7XG4gICAgICAgICAgICAgICAgICAgICAgYm94LXNpemluZzogYm9yZGVyLWJveDtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICBbZGF0YS10ZXJtaW5hbF06YmVmb3JlIHtcbiAgICAgICAgICAgICAgY29udGVudDogJyc7XG4gICAgICAgICAgICAgIHBvc2l0aW9uOiBhYnNvbHV0ZTtcbiAgICAgICAgICAgICAgdG9wOiAxNXB4O1xuICAgICAgICAgICAgICBsZWZ0OiAxNXB4O1xuICAgICAgICAgICAgICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7XG4gICAgICAgICAgICAgIHdpZHRoOiAxNXB4O1xuICAgICAgICAgICAgICBoZWlnaHQ6IDE1cHg7XG4gICAgICAgICAgICAgIGJvcmRlci1yYWRpdXM6IDUwJTtcbiAgICAgICAgICAgICAgLyogQSBsaXR0bGUgaGFjayB0byBkaXNwbGF5IHRoZSB3aW5kb3cgYnV0dG9ucyBpbiBvbmUgcHNldWRvIGVsZW1lbnQuICovXG4gICAgICAgICAgICAgIGJhY2tncm91bmQ6ICNkOTUxNWQ7XG4gICAgICAgICAgICAgIC13ZWJraXQtYm94LXNoYWRvdzogMjVweCAwIDAgI2Y0YzAyNSwgNTBweCAwIDAgIzNlYzkzMDtcbiAgICAgICAgICAgICAgICAgICAgICBib3gtc2hhZG93OiAyNXB4IDAgMCAjZjRjMDI1LCA1MHB4IDAgMCAjM2VjOTMwO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC5jdGEtZm9vdGVyIHtcbiAgICAgICAgICAgIGRpc3BsYXk6IGdyaWQ7XG4gICAgICAgICAgICBncmlkLXRlbXBsYXRlLWNvbHVtbnM6IC43NWZyIDEuMjVmcjtcbiAgICAgICAgICAgIGp1c3RpZnktaXRlbXM6IGNlbnRlcjtcbiAgICAgICAgICAgIGFsaWduLWl0ZW1zOiBjZW50ZXI7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLmN0YS1mb290ZXIgaDIge1xuICAgICAgICAgICAgdGV4dC1hbGlnbjogbGVmdCAhaW1wb3J0YW50O1xuICAgICAgICAgICAgbWFyZ2luOiAwO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC5jdGEtZm9vdGVyIC5idG46Zmlyc3QtY2hpbGQge1xuICAgICAgICAgICAgbWFyZ2luLXJpZ2h0OiAyMHB4O1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC5jdGEtZm9vdGVyIC5idG4ge1xuICAgICAgICAgICAgYm94LXNoYWRvdzogMCAyMHB4IDUwcHggMCByZ2JhKDAsMCwwLDAuMik7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLmhlcm8gLmJ0biB7XG4gICAgICAgICAgICBib3gtc2hhZG93OiAwIDIwcHggNTBweCAwIHJnYmEoMCwwLDAsMC4yKTtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAudGhyZWUge1xuICAgICAgICAgICAgZGlzcGxheTogZ3JpZDtcbiAgICAgICAgICAgIGdyaWQtdGVtcGxhdGUtY29sdW1uczogMWZyIDFmciAxZnI7XG4gICAgICAgICAgICBjb2x1bW4tZ2FwOiAxLjVyZW07XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLnRocmVlIGRpdiB7XG4gICAgICAgICAgICBqdXN0aWZ5LXNlbGY6IGNlbnRlcjtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAudGFiLXNlY3Rpb24ge1xuICAgICAgICAgICAgZGlzcGxheTogZ3JpZDtcbiAgICAgICAgICAgIGdyaWQtdGVtcGxhdGUtY29sdW1uczogMWZyO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC50aHJlZSBoMiB7XG4gICAgICAgICAgICBtYXJnaW46IDA7XG4gICAgICAgICAgICB0ZXh0LWFsaWduOiBsZWZ0ICFpbXBvcnRhbnQ7XG4gICAgICAgICAgICBmb250LXNpemU6IDEwMCUgIWltcG9ydGFudDtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAudGFiLXNlY3Rpb24gdWwge1xuICAgICAgICAgICAganVzdGlmeS1zZWxmOiBjZW50ZXI7XG4gICAgICAgICAgICBsaXN0LXN0eWxlOiBub25lO1xuICAgICAgICAgICAgbWFyZ2luOiAwO1xuICAgICAgICAgICAgcGFkZGluZzogMDtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAudGFiLXNlY3Rpb24gbGkge1xuICAgICAgICAgICAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xuICAgICAgICAgICAgZmxvYXQ6IGxlZnQ7XG4gICAgICAgICAgICBwYWRkaW5nOiA1cHggMjBweDtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAudGFiLXNlY3Rpb24gbGkuYWN0aXZlIHtcbiAgICAgICAgICAgIGJvcmRlci1ib3R0b206IDFweCBzb2xpZCAjMTIyMjM5O1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC50YWNiLXNlY3Rpb24gbGk6aG92ZXIge1xuICAgICAgICAgICAgY3Vyc29yOiBwb2ludGVyO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC50YWItc2VjdGlvbiBidXR0b24ge1xuICAgICAgICAgICAgb3V0bGluZTogbm9uZTtcbiAgICAgICAgICAgIGJvcmRlcjogbm9uZTtcbiAgICAgICAgICAgIGNvbG9yOiAjMTIyMjM5O1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC50YWItc2VjdGlvbiBidXR0b246aG92ZXIge1xuICAgICAgICAgICAgY29sb3I6ICMxMjIyMzk7XG4gICAgICAgICAgICBjdXJzb3I6IHBvaW50ZXI7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLnRhYi1zZWN0aW9uIGJ1dHRvbiB7XG4gICAgICAgICAgICBmb250LXNpemU6IDkwJSAhaW1wb3J0YW50O1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC50d28tc2lkZWQge1xuICAgICAgICAgICAgbWFyZ2luOiAzcmVtIDA7XG4gICAgICAgICAgICBkaXNwbGF5OiBncmlkO1xuICAgICAgICAgICAgZ3JpZC10ZW1wbGF0ZS1jb2x1bW5zOiAxZnIgMWZyO1xuICAgICAgICAgICAgY29sdW1uLWdhcDogMXJlbTtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAudHdvLXNpZGVkIC5wbGFjZWhvbGRlciB7XG4gICAgICAgICAgICBoZWlnaHQ6IDEwMCU7XG4gICAgICAgICAgICB3aWR0aDogMTAwJTtcbiAgICAgICAgICAgIGJhY2tncm91bmQ6ICMxMjIyMzk7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLnR3by1zaWRlZCBwIHtcbiAgICAgICAgICAgIG1hcmdpbjogMXJlbSAwO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC5oZXJvIHtcbiAgICAgICAgICAgIHotaW5kZXg6IDk5OTtcbiAgICAgICAgICAgIHBhZGRpbmc6IDEwMHB4IDAgMTUwcHggMDtcbiAgICAgICAgICAgIHdpZHRoOiA5MCU7XG4gICAgICAgICAgICBwb3NpdGlvbjogcmVsYXRpdmU7XG4gICAgICAgICAgICBsZWZ0OiA1JTtcbiAgICAgICAgICAgIGZvbnQtc2l6ZTogMTMyJTtcbiAgICAgICAgICAgIHRleHQtYWxpZ246IGNlbnRlcjtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAuaGVybyBoMiB7XG4gICAgICAgICAgICBmb250LXNpemU6IDIuMTIxcmVtO1xuICAgICAgICAgICAgbWFyZ2luLWJvdHRvbTogMS41cmVtO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC5oZXJvLWFjdGlvbnMge1xuICAgICAgICAgICAgZm9udC1zaXplOiA4NCU7XG4gICAgICAgICAgICBtYXJnaW46IDQwcHggMCAwIDA7XG4gICAgICAgICAgICBkaXNwbGF5OiBncmlkO1xuICAgICAgICAgICAgY29sdW1uLWdhcDogMXJlbTtcbiAgICAgICAgICAgIGdyaWQtdGVtcGxhdGUtY29sdW1uczogMWZyIDFmcjtcbiAgICAgICAgICAgIGFsaWduLWl0ZW1zOiBjZW50ZXI7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLmhlcm8tYWN0aW9ucyBkaXY6Zmlyc3QtY2hpbGQge1xuICAgICAgICAgICAganVzdGlmeS1zZWxmOiBlbmQ7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLmhlcm8tYWN0aW9ucyBkaXY6bGFzdC1jaGlsZCB7XG4gICAgICAgICAgICBqdXN0aWZ5LXNlbGY6IHN0YXJ0O1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC5zaG93Y2FzZSB7XG4gICAgICAgICAgICBiYWNrZ3JvdW5kOiAjZmFmYmZjO1xuICAgICAgICAgICAgcGFkZGluZzogM3JlbSAwO1xuICAgICAgICAgICAgYm9yZGVyLWJvdHRvbTogMXB4IHNvbGlkICNlYWVhZWE7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLnNob3djYXNlIGgyIHtcbiAgICAgICAgICAgIHRleHQtYWxpZ246IGNlbnRlcjtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAuc2hvd2Nhc2UuYWx0IHtcbiAgICAgICAgICAgIGJhY2tncm91bmQ6ICNmZmY7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLnNob3djYXNlLmFsdCBwIHtcbiAgICAgICAgICAgIGNvbG9yOiAjNzA3ZDhlO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC5mdWxsLXdpZHRoIHtcbiAgICAgICAgICAgIHdpZHRoOiAxMDB2dztcbiAgICAgICAgICAgIHBvc2l0aW9uOiByZWxhdGl2ZTtcbiAgICAgICAgICAgIGxlZnQ6IDUwJTtcbiAgICAgICAgICAgIHJpZ2h0OiA1MCU7XG4gICAgICAgICAgICBtYXJnaW4tbGVmdDogLTUwdnc7XG4gICAgICAgICAgICBtYXJnaW4tcmlnaHQ6IC01MHZ3O1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC5sb25nLXdpZHRoIHtcbiAgICAgICAgICAgIHBvc2l0aW9uOiByZWxhdGl2ZTtcbiAgICAgICAgICAgIHdpZHRoOiBjYWxjKDEwMCUgKyAxNDZweCk7XG4gICAgICAgICAgICBtYXJnaW4tbGVmdDogLTczcHg7XG4gICAgICAgICAgICBtYXJnaW4tcmlnaHQ6IC03M3B4O1xuICAgICAgICAgIH1cblxuICAgICAgICAgIC50ZXJtaW5hbCB7XG4gICAgICAgICAgICBtYXJnaW46IDAgYXV0bztcbiAgICAgICAgICAgIGJhY2tncm91bmQ6ICMxMjIyMzk7XG4gICAgICAgICAgICBib3gtc2hhZG93OiAwIDIwcHggNTBweCAwIHJnYmEoMCwwLDAsMC4yKTtcbiAgICAgICAgICAgIGJvcmRlci1yYWRpdXM6IDVweDtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAudGVybWluYWwgaW1nIHtcbiAgICAgICAgICAgIHdpZHRoOiAxMDAlO1xuICAgICAgICAgIH1cbiAgICAgICAgYH08L3N0eWxlPlxuICAgICAgPC8+XG4gICAgKVxuICB9XG59XG5cbi8vIGZpbHRlciByZWR1eCBzdGF0ZSB0byBwcm9wcyBmb3IgdGhpcyBwYWdlXG5jb25zdCBtYXBTdGF0ZVRvUHJvcHMgPSAoeyBwb3N0cyB9KSA9PiAoeyBwb3N0cyB9KVxuXG4vLyBjb25uZWN0IHJlZHV4IHN0YXRlIHRvIGNvbXBvbmVudFxuZXhwb3J0IGRlZmF1bHQgY29ubmVjdChcbiAgbWFwU3RhdGVUb1Byb3BzLFxuICBudWxsXG4pKEhvbWUpIl19 */\n/*@ sourceURL=/Users/dev/go/src/github.com/dannav/starship/web/components/homepage.js */"
      }));
    }
  }]);

  return Home;
}(react__WEBPACK_IMPORTED_MODULE_1___default.a.Component); // filter redux state to props for this page


var mapStateToProps = function mapStateToProps(_ref) {
  var posts = _ref.posts;
  return {
    posts: posts
  };
}; // connect redux state to component


/* harmony default export */ __webpack_exports__["default"] = (Object(react_redux__WEBPACK_IMPORTED_MODULE_2__["connect"])(mapStateToProps, null)(Home));

/***/ }),

/***/ "./components/shared/layout.js":
/*!*************************************!*\
  !*** ./components/shared/layout.js ***!
  \*************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var styled_jsx_style__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! styled-jsx/style */ "styled-jsx/style");
/* harmony import */ var styled_jsx_style__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(styled_jsx_style__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ "react");
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var react_redux__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! react-redux */ "react-redux");
/* harmony import */ var react_redux__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(react_redux__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! next/link */ "next/link");
/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(next_link__WEBPACK_IMPORTED_MODULE_3__);





var childrenWithProps = function childrenWithProps(children, props) {
  return react__WEBPACK_IMPORTED_MODULE_1___default.a.Children.map(children, function (child) {
    return react__WEBPACK_IMPORTED_MODULE_1___default.a.cloneElement(child, props);
  });
}; // connect redux state to component


/* harmony default export */ __webpack_exports__["default"] = (Object(react_redux__WEBPACK_IMPORTED_MODULE_2__["connect"])(function (state) {
  return state;
})(function (props) {
  return react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(react__WEBPACK_IMPORTED_MODULE_1___default.a.Fragment, null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("header", {
    className: "container"
  }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("h1", null, props.title), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("nav", {
    className: "menu"
  }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("ul", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("li", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(next_link__WEBPACK_IMPORTED_MODULE_3___default.a, {
    prefetch: true,
    href: "/"
  }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("a", null, "Home"))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("li", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(next_link__WEBPACK_IMPORTED_MODULE_3___default.a, {
    prefetch: true,
    href: "/contact"
  }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("a", null, "Use Cases"))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("li", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(next_link__WEBPACK_IMPORTED_MODULE_3___default.a, {
    prefetch: true,
    href: "/guide"
  }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("a", null, "Documentation"))))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("nav", {
    className: "header-actions"
  }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("ul", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("li", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(next_link__WEBPACK_IMPORTED_MODULE_3___default.a, {
    prefetch: true,
    href: "/"
  }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("a", {
    className: "dwnld"
  }, "Download"))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("li", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(next_link__WEBPACK_IMPORTED_MODULE_3___default.a, {
    prefetch: true,
    href: "/contact"
  }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("a", {
    className: "btn"
  }, "Login / Sign up")))))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("main", null, childrenWithProps(props.children, props)), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("footer", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("div", {
    className: "container"
  }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("nav", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("ul", null, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("li", null, "Privacy & Terms"), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement("li", null, "Contact Us"))))), react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(styled_jsx_style__WEBPACK_IMPORTED_MODULE_0___default.a, {
    styleId: "1866397126",
    css: "footer{font-weight:bold;color:#122239;padding:1.5rem 0;}footer ul{list-style:none;margin:0;padding:0;}footer li{display:inline-block;font-size:80%;margin-right:1rem;}header.container{max-width:1090px;}.container{max-width:940px;margin:0 auto;}.btn{padding:8px 20px;text-align:center;border-radius:7px;background:#122239;color:#fff !important;font-size:90%;border:none !important;}header .btn{padding:5px 15px;}.btn:hover{border:none;}header{width:100%;display:grid;grid-template-columns:.75fr 1fr .75fr;-webkit-align-items:center;-webkit-box-align:center;-ms-flex-align:center;align-items:center;height:100px;box-sizing:border-box;}header h1{margin:0;font-size:1.414rem;position:relative;top:-2px;}header .menu{justify-self:center;}header .header-actions{justify-self:end;}header .dwnld{color:#707d8e !important;}header .dwnld:hover{border:0;color:#122239 !important;}header a,header a:hover,header a:visited,header a:focus{color:#122239;-webkit-text-decoration:none;text-decoration:none;}header a:visited{color:black;}header a:hover{border-bottom:1px solid #122239;}header ul{padding:0;margin:0;list-style:none;}header li{display:inline-block;padding-left:20px;}header li:first-child{padding:0;}\n/*# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIi9Vc2Vycy9kZXYvZ28vc3JjL2dpdGh1Yi5jb20vZGFubmF2L3N0YXJzaGlwL3dlYi9jb21wb25lbnRzL3NoYXJlZC9sYXlvdXQuanMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IkFBK0QyQixBQUc4QixBQU1ELEFBTUssQUFNSixBQUlELEFBS0MsQUFVQSxBQUlMLEFBSUQsQUFTRixBQU9XLEFBSUgsQUFJUSxBQUloQixBQUtLLEFBS0YsQUFJb0IsQUFJdEIsQUFNVyxBQUtYLFNBL0NTLEFBbUJNLENBa0JoQixBQVdYLENBekRlLENBSmYsQUEwQ0EsRUFMdUIsRUF4RVosQUFnQkssQ0F0QkEsQUFrQmhCLEFBU29CLEFBVXBCLEFBNEJBLEVBMkJrQixDQS9CbEIsQ0FqRGdCLEFBcUZJLEdBbkRvQixDQXZDNUIsQUE4RFosR0Fkb0IsRUFoQ3BCLENBdEJtQixDQXNGbkIsRUFiQSxDQWxFQSxBQUtvQixBQWVBLEFBaUVwQixJQUtBLE9BMUNXLEVBdERYLEtBWUEsQUFlcUIsRUE0QnJCLE9BVnFCLEVBZ0NyQixRQWpEd0Isc0JBQ1IsY0FDUyx1QkFDekIsd0JBZWUsYUFDUyxzQkFDeEIiLCJmaWxlIjoiL1VzZXJzL2Rldi9nby9zcmMvZ2l0aHViLmNvbS9kYW5uYXYvc3RhcnNoaXAvd2ViL2NvbXBvbmVudHMvc2hhcmVkL2xheW91dC5qcyIsInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IGNvbm5lY3QgfSBmcm9tICdyZWFjdC1yZWR1eCdcbmltcG9ydCBMaW5rIGZyb20gJ25leHQvbGluaydcblxuY29uc3QgY2hpbGRyZW5XaXRoUHJvcHMgPSAoY2hpbGRyZW4sIHByb3BzKSA9PiBSZWFjdC5DaGlsZHJlbi5tYXAoY2hpbGRyZW4sIGNoaWxkID0+XG4gIFJlYWN0LmNsb25lRWxlbWVudChjaGlsZCwgcHJvcHMpXG4pO1xuXG4vLyBjb25uZWN0IHJlZHV4IHN0YXRlIHRvIGNvbXBvbmVudFxuZXhwb3J0IGRlZmF1bHQgY29ubmVjdChzdGF0ZSA9PiBzdGF0ZSkoXG4gIHByb3BzID0+IHtcbiAgICByZXR1cm4gKFxuICAgICAgPD5cbiAgICAgICAgPGhlYWRlciBjbGFzc05hbWU9XCJjb250YWluZXJcIj5cbiAgICAgICAgICA8aDE+e3Byb3BzLnRpdGxlfTwvaDE+XG4gICAgICAgICAgPG5hdiBjbGFzc05hbWU9XCJtZW51XCI+XG4gICAgICAgICAgICA8dWw+XG4gICAgICAgICAgICAgIDxsaT5cbiAgICAgICAgICAgICAgICA8TGluayBwcmVmZXRjaCBocmVmPVwiL1wiPlxuICAgICAgICAgICAgICAgICAgPGE+SG9tZTwvYT5cbiAgICAgICAgICAgICAgICA8L0xpbms+XG4gICAgICAgICAgICAgIDwvbGk+XG4gICAgICAgICAgICAgIDxsaT5cbiAgICAgICAgICAgICAgICA8TGluayBwcmVmZXRjaCBocmVmPVwiL2NvbnRhY3RcIj5cbiAgICAgICAgICAgICAgICAgIDxhPlVzZSBDYXNlczwvYT5cbiAgICAgICAgICAgICAgICA8L0xpbms+XG4gICAgICAgICAgICAgIDwvbGk+XG4gICAgICAgICAgICAgIDxsaT5cbiAgICAgICAgICAgICAgICA8TGluayBwcmVmZXRjaCBocmVmPVwiL2d1aWRlXCI+XG4gICAgICAgICAgICAgICAgICA8YT5Eb2N1bWVudGF0aW9uPC9hPlxuICAgICAgICAgICAgICAgIDwvTGluaz5cbiAgICAgICAgICAgICAgPC9saT5cbiAgICAgICAgICAgIDwvdWw+XG4gICAgICAgICAgPC9uYXY+XG4gICAgICAgICAgPG5hdiBjbGFzc05hbWU9XCJoZWFkZXItYWN0aW9uc1wiPlxuICAgICAgICAgICAgPHVsPlxuICAgICAgICAgICAgICA8bGk+XG4gICAgICAgICAgICAgICAgPExpbmsgcHJlZmV0Y2ggaHJlZj1cIi9cIj5cbiAgICAgICAgICAgICAgICAgIDxhIGNsYXNzTmFtZT1cImR3bmxkXCI+RG93bmxvYWQ8L2E+XG4gICAgICAgICAgICAgICAgPC9MaW5rPlxuICAgICAgICAgICAgICA8L2xpPlxuICAgICAgICAgICAgICA8bGk+XG4gICAgICAgICAgICAgICAgPExpbmsgcHJlZmV0Y2ggaHJlZj1cIi9jb250YWN0XCI+XG4gICAgICAgICAgICAgICAgICA8YSBjbGFzc05hbWU9XCJidG5cIj5Mb2dpbiAvIFNpZ24gdXA8L2E+XG4gICAgICAgICAgICAgICAgPC9MaW5rPlxuICAgICAgICAgICAgICA8L2xpPlxuICAgICAgICAgICAgPC91bD5cbiAgICAgICAgICA8L25hdj5cbiAgICAgICAgPC9oZWFkZXI+XG4gICAgICAgIDxtYWluPlxuICAgICAgICAgIHtcbiAgICAgICAgICAgIGNoaWxkcmVuV2l0aFByb3BzKHByb3BzLmNoaWxkcmVuLCBwcm9wcylcbiAgICAgICAgICB9XG4gICAgICAgIDwvbWFpbj5cbiAgICAgICAgPGZvb3Rlcj5cbiAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cImNvbnRhaW5lclwiPlxuICAgICAgICAgICAgPG5hdj5cbiAgICAgICAgICAgICAgPHVsPlxuICAgICAgICAgICAgICAgIDxsaT5Qcml2YWN5ICZhbXA7IFRlcm1zPC9saT5cbiAgICAgICAgICAgICAgICA8bGk+Q29udGFjdCBVczwvbGk+XG4gICAgICAgICAgICAgIDwvdWw+XG4gICAgICAgICAgICA8L25hdj5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgPC9mb290ZXI+XG4gICAgICAgIDxzdHlsZSBnbG9iYWwganN4PntgXG4gICAgICAgICAgZm9vdGVyIHtcbiAgICAgICAgICAgIGZvbnQtd2VpZ2h0OiBib2xkO1xuICAgICAgICAgICAgY29sb3I6ICMxMjIyMzk7XG4gICAgICAgICAgICBwYWRkaW5nOiAxLjVyZW0gMDtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICBmb290ZXIgdWwge1xuICAgICAgICAgICAgbGlzdC1zdHlsZTogbm9uZTtcbiAgICAgICAgICAgIG1hcmdpbjogMDtcbiAgICAgICAgICAgIHBhZGRpbmc6IDA7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgZm9vdGVyIGxpIHtcbiAgICAgICAgICAgIGRpc3BsYXk6IGlubGluZS1ibG9jaztcbiAgICAgICAgICAgIGZvbnQtc2l6ZTogODAlO1xuICAgICAgICAgICAgbWFyZ2luLXJpZ2h0OiAxcmVtO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIGhlYWRlci5jb250YWluZXIge1xuICAgICAgICAgICAgbWF4LXdpZHRoOiAxMDkwcHg7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLmNvbnRhaW5lciB7XG4gICAgICAgICAgICBtYXgtd2lkdGg6IDk0MHB4O1xuICAgICAgICAgICAgbWFyZ2luOiAwIGF1dG87XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLmJ0biB7XG4gICAgICAgICAgICBwYWRkaW5nOiA4cHggMjBweDtcbiAgICAgICAgICAgIHRleHQtYWxpZ246IGNlbnRlcjtcbiAgICAgICAgICAgIGJvcmRlci1yYWRpdXM6IDdweDtcbiAgICAgICAgICAgIGJhY2tncm91bmQ6ICMxMjIyMzk7XG4gICAgICAgICAgICBjb2xvcjogI2ZmZiAhaW1wb3J0YW50O1xuICAgICAgICAgICAgZm9udC1zaXplOiA5MCU7XG4gICAgICAgICAgICBib3JkZXI6IG5vbmUgIWltcG9ydGFudDtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICBoZWFkZXIgLmJ0biB7XG4gICAgICAgICAgICBwYWRkaW5nOiA1cHggMTVweDtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICAuYnRuOmhvdmVyIHtcbiAgICAgICAgICAgIGJvcmRlcjogbm9uZTtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICBoZWFkZXIge1xuICAgICAgICAgICAgd2lkdGg6IDEwMCU7XG4gICAgICAgICAgICBkaXNwbGF5OiBncmlkO1xuICAgICAgICAgICAgZ3JpZC10ZW1wbGF0ZS1jb2x1bW5zOiAuNzVmciAxZnIgLjc1ZnI7XG4gICAgICAgICAgICBhbGlnbi1pdGVtczogY2VudGVyO1xuICAgICAgICAgICAgaGVpZ2h0OiAxMDBweDtcbiAgICAgICAgICAgIGJveC1zaXppbmc6IGJvcmRlci1ib3g7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgaGVhZGVyIGgxIHtcbiAgICAgICAgICAgIG1hcmdpbjogMDtcbiAgICAgICAgICAgIGZvbnQtc2l6ZTogMS40MTRyZW07XG4gICAgICAgICAgICBwb3NpdGlvbjogcmVsYXRpdmU7XG4gICAgICAgICAgICB0b3A6IC0ycHg7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgaGVhZGVyIC5tZW51IHtcbiAgICAgICAgICAgIGp1c3RpZnktc2VsZjogY2VudGVyO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIGhlYWRlciAuaGVhZGVyLWFjdGlvbnMge1xuICAgICAgICAgICAganVzdGlmeS1zZWxmOiBlbmQ7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgaGVhZGVyIC5kd25sZCB7XG4gICAgICAgICAgICBjb2xvcjogIzcwN2Q4ZSAhaW1wb3J0YW50O1xuICAgICAgICAgIH1cblxuICAgICAgICAgIGhlYWRlciAuZHdubGQ6aG92ZXIge1xuICAgICAgICAgICAgYm9yZGVyOiAwO1xuICAgICAgICAgICAgY29sb3I6ICMxMjIyMzkgIWltcG9ydGFudDtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICBoZWFkZXIgYSwgaGVhZGVyIGE6aG92ZXIsIGhlYWRlciBhOnZpc2l0ZWQsIGhlYWRlciBhOmZvY3VzIHtcbiAgICAgICAgICAgIGNvbG9yOiAjMTIyMjM5O1xuICAgICAgICAgICAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIGhlYWRlciBhOnZpc2l0ZWQge1xuICAgICAgICAgICAgY29sb3I6IGJsYWNrO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIGhlYWRlciBhOmhvdmVyIHtcbiAgICAgICAgICAgIGJvcmRlci1ib3R0b206IDFweCBzb2xpZCAjMTIyMjM5O1xuICAgICAgICAgIH1cblxuICAgICAgICAgIGhlYWRlciB1bCB7XG4gICAgICAgICAgICBwYWRkaW5nOiAwO1xuICAgICAgICAgICAgbWFyZ2luOiAwO1xuICAgICAgICAgICAgbGlzdC1zdHlsZTogbm9uZTtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICBoZWFkZXIgbGkge1xuICAgICAgICAgICAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xuICAgICAgICAgICAgcGFkZGluZy1sZWZ0OiAyMHB4O1xuICAgICAgICAgIH1cblxuICAgICAgICAgIGhlYWRlciBsaTpmaXJzdC1jaGlsZCB7XG4gICAgICAgICAgICBwYWRkaW5nOiAwO1xuICAgICAgICAgIH1cbiAgICAgICAgYH08L3N0eWxlPlxuICAgICAgPC8+XG4gICAgKVxuICB9XG4pIl19 */\n/*@ sourceURL=/Users/dev/go/src/github.com/dannav/starship/web/components/shared/layout.js */"
  }));
}));

/***/ }),

/***/ "./pages/index.js":
/*!************************!*\
  !*** ./pages/index.js ***!
  \************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @babel/runtime/regenerator */ "@babel/runtime/regenerator");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ "react");
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var react_redux__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! react-redux */ "react-redux");
/* harmony import */ var react_redux__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(react_redux__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _components_homepage__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../components/homepage */ "./components/homepage.js");
/* harmony import */ var _components_shared_layout__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../components/shared/layout */ "./components/shared/layout.js");
/* harmony import */ var _redux_store__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ../redux/store */ "./redux/store.js");


function _typeof(obj) { if (typeof Symbol === "function" && typeof Symbol.iterator === "symbol") { _typeof = function _typeof(obj) { return typeof obj; }; } else { _typeof = function _typeof(obj) { return obj && typeof Symbol === "function" && obj.constructor === Symbol && obj !== Symbol.prototype ? "symbol" : typeof obj; }; } return _typeof(obj); }

function asyncGeneratorStep(gen, resolve, reject, _next, _throw, key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { Promise.resolve(value).then(_next, _throw); } }

function _asyncToGenerator(fn) { return function () { var self = this, args = arguments; return new Promise(function (resolve, reject) { var gen = fn.apply(self, args); function _next(value) { asyncGeneratorStep(gen, resolve, reject, _next, _throw, "next", value); } function _throw(err) { asyncGeneratorStep(gen, resolve, reject, _next, _throw, "throw", err); } _next(undefined); }); }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); return Constructor; }

function _possibleConstructorReturn(self, call) { if (call && (_typeof(call) === "object" || typeof call === "function")) { return call; } return _assertThisInitialized(self); }

function _assertThisInitialized(self) { if (self === void 0) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return self; }

function _getPrototypeOf(o) { _getPrototypeOf = Object.setPrototypeOf ? Object.getPrototypeOf : function _getPrototypeOf(o) { return o.__proto__ || Object.getPrototypeOf(o); }; return _getPrototypeOf(o); }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function"); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, writable: true, configurable: true } }); if (superClass) _setPrototypeOf(subClass, superClass); }

function _setPrototypeOf(o, p) { _setPrototypeOf = Object.setPrototypeOf || function _setPrototypeOf(o, p) { o.__proto__ = p; return o; }; return _setPrototypeOf(o, p); }







var Homepage =
/*#__PURE__*/
function (_React$Component) {
  _inherits(Homepage, _React$Component);

  function Homepage() {
    _classCallCheck(this, Homepage);

    return _possibleConstructorReturn(this, _getPrototypeOf(Homepage).apply(this, arguments));
  }

  _createClass(Homepage, [{
    key: "render",
    value: function render() {
      return react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(_components_shared_layout__WEBPACK_IMPORTED_MODULE_4__["default"], {
        title: "Starship"
      }, react__WEBPACK_IMPORTED_MODULE_1___default.a.createElement(_components_homepage__WEBPACK_IMPORTED_MODULE_3__["default"], null));
    }
  }], [{
    key: "getInitialProps",
    value: function () {
      var _getInitialProps = _asyncToGenerator(
      /*#__PURE__*/
      _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0___default.a.mark(function _callee(_ref) {
        var store, isServer;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0___default.a.wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                store = _ref.store, isServer = _ref.isServer;
                _context.next = 3;
                return store.dispatch(Object(_redux_store__WEBPACK_IMPORTED_MODULE_5__["getPosts"])(isServer));

              case 3:
                return _context.abrupt("return", {
                  isServer: isServer
                });

              case 4:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, this);
      }));

      function getInitialProps(_x) {
        return _getInitialProps.apply(this, arguments);
      }

      return getInitialProps;
    }()
  }]);

  return Homepage;
}(react__WEBPACK_IMPORTED_MODULE_1___default.a.Component);

/* harmony default export */ __webpack_exports__["default"] = (Object(react_redux__WEBPACK_IMPORTED_MODULE_2__["connect"])(function (state) {
  return state;
})(Homepage));

/***/ }),

/***/ "./redux/store.js":
/*!************************!*\
  !*** ./redux/store.js ***!
  \************************/
/*! exports provided: actionTypes, reducer, getPosts, initStore */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "actionTypes", function() { return actionTypes; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "reducer", function() { return reducer; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "getPosts", function() { return getPosts; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "initStore", function() { return initStore; });
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @babel/runtime/regenerator */ "@babel/runtime/regenerator");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var redux__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! redux */ "redux");
/* harmony import */ var redux__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(redux__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var redux_devtools_extension__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! redux-devtools-extension */ "redux-devtools-extension");
/* harmony import */ var redux_devtools_extension__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(redux_devtools_extension__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var redux_thunk__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! redux-thunk */ "redux-thunk");
/* harmony import */ var redux_thunk__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(redux_thunk__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var isomorphic_unfetch__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! isomorphic-unfetch */ "isomorphic-unfetch");
/* harmony import */ var isomorphic_unfetch__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(isomorphic_unfetch__WEBPACK_IMPORTED_MODULE_4__);


function asyncGeneratorStep(gen, resolve, reject, _next, _throw, key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { Promise.resolve(value).then(_next, _throw); } }

function _asyncToGenerator(fn) { return function () { var self = this, args = arguments; return new Promise(function (resolve, reject) { var gen = fn.apply(self, args); function _next(value) { asyncGeneratorStep(gen, resolve, reject, _next, _throw, "next", value); } function _throw(err) { asyncGeneratorStep(gen, resolve, reject, _next, _throw, "throw", err); } _next(undefined); }); }; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; var ownKeys = Object.keys(source); if (typeof Object.getOwnPropertySymbols === 'function') { ownKeys = ownKeys.concat(Object.getOwnPropertySymbols(source).filter(function (sym) { return Object.getOwnPropertyDescriptor(source, sym).enumerable; })); } ownKeys.forEach(function (key) { _defineProperty(target, key, source[key]); }); } return target; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }





var initialState = {
  posts: []
};
var actionTypes = {
  GET_POSTS: 'GET_POSTS' // REDUCERS

};
var reducer = function reducer() {
  var state = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : initialState;
  var action = arguments.length > 1 ? arguments[1] : undefined;

  switch (action.type) {
    case actionTypes.GET_POSTS:
      return _objectSpread({}, state, {
        posts: action.results
      });

    default:
      return state;
  }
}; // ACTIONS

var getPosts = function getPosts(isServer) {
  return (
    /*#__PURE__*/
    function () {
      var _ref = _asyncToGenerator(
      /*#__PURE__*/
      _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0___default.a.mark(function _callee(dispatch, getState) {
        var _getState, posts, res;

        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0___default.a.wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                _getState = getState(), posts = _getState.posts;

                if (posts.length) {
                  _context.next = 9;
                  break;
                }

                _context.next = 4;
                return isomorphic_unfetch__WEBPACK_IMPORTED_MODULE_4___default()('http://localhost:6060/static/posts.json');

              case 4:
                res = _context.sent;
                _context.next = 7;
                return res.json();

              case 7:
                posts = _context.sent;
                return _context.abrupt("return", dispatch({
                  type: actionTypes.GET_POSTS,
                  results: posts.results
                }));

              case 9:
                return _context.abrupt("return");

              case 10:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, this);
      }));

      return function (_x, _x2) {
        return _ref.apply(this, arguments);
      };
    }()
  );
};
var initStore = function initStore() {
  var initialState = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : initialState;
  return Object(redux__WEBPACK_IMPORTED_MODULE_1__["createStore"])(reducer, initialState, Object(redux_devtools_extension__WEBPACK_IMPORTED_MODULE_2__["composeWithDevTools"])(Object(redux__WEBPACK_IMPORTED_MODULE_1__["applyMiddleware"])(redux_thunk__WEBPACK_IMPORTED_MODULE_3___default.a)));
};

/***/ }),

/***/ 3:
/*!******************************!*\
  !*** multi ./pages/index.js ***!
  \******************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__(/*! ./pages/index.js */"./pages/index.js");


/***/ }),

/***/ "@babel/runtime/regenerator":
/*!*********************************************!*\
  !*** external "@babel/runtime/regenerator" ***!
  \*********************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = require("@babel/runtime/regenerator");

/***/ }),

/***/ "isomorphic-unfetch":
/*!*************************************!*\
  !*** external "isomorphic-unfetch" ***!
  \*************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = require("isomorphic-unfetch");

/***/ }),

/***/ "next/link":
/*!****************************!*\
  !*** external "next/link" ***!
  \****************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = require("next/link");

/***/ }),

/***/ "react":
/*!************************!*\
  !*** external "react" ***!
  \************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = require("react");

/***/ }),

/***/ "react-redux":
/*!******************************!*\
  !*** external "react-redux" ***!
  \******************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = require("react-redux");

/***/ }),

/***/ "redux":
/*!************************!*\
  !*** external "redux" ***!
  \************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = require("redux");

/***/ }),

/***/ "redux-devtools-extension":
/*!*******************************************!*\
  !*** external "redux-devtools-extension" ***!
  \*******************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = require("redux-devtools-extension");

/***/ }),

/***/ "redux-thunk":
/*!******************************!*\
  !*** external "redux-thunk" ***!
  \******************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = require("redux-thunk");

/***/ }),

/***/ "styled-jsx/style":
/*!***********************************!*\
  !*** external "styled-jsx/style" ***!
  \***********************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = require("styled-jsx/style");

/***/ })

/******/ });
//# sourceMappingURL=index.js.map