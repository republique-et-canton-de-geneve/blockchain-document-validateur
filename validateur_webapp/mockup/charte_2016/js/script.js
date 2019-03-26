/**
 * @file
 * A JavaScript file for the theme.
 *
 * In order for this JavaScript to be loaded on pages, see the instructions in
 * the README.txt next to this file.
 */

/* Fonction générique à tous les thèmes */
(function ($) {
  $(document).ready(function(){
    var number = $('.view-prestation .view-content ul li').size();
    $('.view-prestation .view-content ul li:nth-child('+Math.floor(number/2)+')').addClass('lastOfFirst');

    $('.content .col-md-8 a:has(img)').addClass('imgInto');
    $('.content .col-md-9 a:has(img)').addClass('imgInto');

    $("#menu-toggle").click(function(e) {
      e.preventDefault();
      $("#wrapper").toggleClass("toggled");
    });
    $("#close_sidebar").click(function(e) {
      e.preventDefault();
      $("#wrapper").removeClass("toggled");
    });

    // Vérifie la présence du cookie ETATRAD
    if ($.cookie('ETATRAD')) {
      // have cookie
      $("#menu-toggle").addClass("active");

    } else {
      //no cookie
      $("#menu-toggle").removeClass("active");
    }
  });
}(jQuery));


// JavaScript should be made compatible with libraries other than jQuery by
// wrapping it with an "anonymous closure". See:
// - https://drupal.org/node/1446420
// - http://www.adequatelygood.com/2010/3/JavaScript-Module-Pattern-In-Depth
(function ($, Drupal, window, document, undefined) {

  //LISTES
  Drupal.behaviors.Listes = {
    attach: function(context, settings) {
      $('body').ready(function() {
        $( "ol" ).each(function() {
          var   val=1;
          if ( $(this).attr("start")){
            val =  $(this).attr("start");
            val=val-1;
            val= 'list-body '+ val;
            $(this ).css('counter-increment',val );
          }

        });
      });
    }
  };


  //RESPONSIVE TABLE
  Drupal.behaviors.Table = {
    attach: function(context, settings) {
      $('body').ready(function() {
        $('table').wrap('<div class="table-responsive"></div>');
      });
    }
  };

  //cf CHARTEDRUPAL-323
  Drupal.behaviors.LinkImage = {
    attach: function(context, settings) {
      $('body').ready(function() {
        $( "a img" ).each(function() {
          $(this).parent().addClass('liens-images-externes');
        });
      });
    }
  };


  //OWL CAROUSEL
  Drupal.behaviors.Carousel = {
    attach: function(context, settings) {
      $('.field-type-image .field-items, .views-field-field-images .field-content').owlCarousel({
        navigation : true,
        slideSpeed : 300,
        paginationSpeed : 400,
        singleItem:true
      })
    }
  };

  Drupal.behaviors.changeHash = {
    attach: function (context, settings) {
      $(window).on('hashchange', function (e) {
        e.preventDefault();
        return false;
      });
    }
  };

  // REMARQUE SUR LES PRESTATIONS
  /*Drupal.behaviors.sidebar = {
    attach: function(context, settings) {
      $("#menu-toggle").click(function(e) {
        e.preventDefault();
        $("#wrapper").toggleClass("toggled");
      });
      $("#close_sidebar").click(function(e) {
        e.preventDefault();
        $("#wrapper").removeClass("toggled");
      });

    }
  }*/

  //IMAGE CAPTION
  /*
  Drupal.behaviors.caption = {
    attach: function(context, settings) {
      $('body').ready(function() {
        $(".owl-theme").hover(function() {
          $(this).find('figure').trigger('mouseover');
        }, function() {
          $(this).find('figure').trigger('mouseout');
        });
      });

      $('figure').hover(function(event) {
        $('figure figcaption').css({"opacity":"0.5", "bottom": "0"});
        event.stopPropagation();
      }, function(event) {
        $('figure figcaption').css({"opacity":"0", "bottom": "-25px"});
        event.stopPropagation();
      });
    }
  };
  */

  // REMARQUE SUR LES PRESTATIONS
  Drupal.behaviors.edgNotice = {
    attach: function(context, settings) {
      $(".pane-edg-2016-base-comment-pane .pane-title").click(function() {
        $(this).siblings().children('form').toggle(1000);
      });
    }
  }  

  //MENU
  Drupal.behaviors.Menu = {
    attach: function (context, settings) {
      var MenuItem = function (menu) {
        var self = this;
        self.menu = menu;
        if (self.menu.attr("class") == 'active') {
          activeMenu = self;
        }
        self.name = self.menu.attr("href").substring(1);
        self.items = [];
        self.submenu = level2.find(".menu-category" + self.menu.attr("href"));
        self.submenu.find("li a").each(function() {
          self.items.push(this);
          if ($(this).attr("class") == 'active') {
            activeSubMenu = $(this);
          }
        });

        self.menu.click(function(e) {
          e.preventDefault();
          replaceActiveMenu(self, null);
          if(history.pushState) {
            history.pushState(null, null, "#" + self.name);
          }
          else {
            window.location.hash = "#" + self.name;
          }
          return false;
        });
        if (typeof self.items == "object") {
          self.items.map(function(item) {
            $(item).click(function(e) {
              e.preventDefault();
              replaceActiveSubMenu($(this));
              var stateName = self.name + "-" + $(this).attr("href").substring(1);
              if(history.pushState) {
                history.pushState(null, null, "#" + stateName);
              }
              else {
                window.location.hash = "#" + stateName;
              }
              return false;
            });
          });
        }
      };

      var MenuStorage = [];

      var activeMenu = null;
      var activeSubMenu = null;

      function replaceActiveSubMenu(submenu) {
        if (activeSubMenu !== null) {
          activeSubMenu.removeClass("active");
          level3.find(".menu-category" + activeSubMenu.attr("href")).removeClass("active");
        }
        if (submenu !== null) {
          submenu.addClass("active");
          level3.find(".menu-category" + submenu.attr("href")).addClass("active");
          setBrowseStep("browse-step-2");
          $("#browse-back").html(submenu.find("h3").html());
        }
        activeSubMenu = submenu;
        setTimeout(function(){
          window.scrollTo(0, 0);
        }, 50);
      }

      function replaceActiveMenu(menu, submenu) {
        if (activeMenu !== null) {
          activeMenu.menu.removeClass("active");
          activeMenu.submenu.removeClass("active");
        }
        if (menu !== null) {
          menu.menu.addClass("active");
          menu.submenu.addClass("active");
          setBrowseStep("browse-step-1");
          $("#browse-back").html(menu.menu.html());
        }
        else {
          setBrowseStep(null);
        }
        activeMenu = menu;
        replaceActiveSubMenu(submenu);
      }

      function initHash(hash) {
        if (hash.length > 1) {
          var names = hash.split("-");
          MenuStorage.map(function(item){
            if (item.name == names[0]) {
              submenu = null;
              if (typeof names[1] !== "undefined") {
                submenu = $("a[href=#" + names[1] + "]");
              }
              replaceActiveMenu(item, submenu);
            }
          });
          $('html, body').animate({ scrollTop: 0 }, 0);
        }
        else {
          replaceActiveMenu(null, null);
        }
      }

      function setBrowseStep(step) {
        $('.browse-page').removeClass('browse-step-1');
        $('.browse-page').removeClass('browse-step-2');
        $('.browse-level').removeClass('active');
        if (step !== null) {
          $('.browse-page').addClass(step);
          $('.browse-page').removeClass('browse-level-1');

          if (step == 'browse-step-1') {
            level2.addClass('active');
          }
          if (step == 'browse-step-2') {
            level3.addClass('active');
          }
        }
        else {
          level1.addClass('active');
          $('.browse-page').addClass('browse-level-1');
        }
      }

      var level1 = $(".browse-page .browse-level-1");
      var level2 = $(".browse-page .browse-level-2");
      var level3 = $(".browse-page .browse-level-3");

      level1.find("li a").each(function() {
        MenuStorage.push(new MenuItem($(this)));
      });

      initHash(History.getHash());

      $("#browse-back").click(function(e){
        if (activeSubMenu !== null) {
          replaceActiveMenu(activeMenu, null);
          var stateName = activeMenu.name;
        }
        else {
          replaceActiveMenu(null, null);
          var stateName = '';
        }
        e.preventDefault();
        if(history.pushState) {
          history.pushState(null, null, "#" + stateName);
        }
        else {
          window.location.hash = "#" + stateName;
        }
        //$('html, body').animate({ scrollTop: 0 }, 0);
        return false;
      });
    }
  };

})(jQuery, Drupal, this, this.document);
