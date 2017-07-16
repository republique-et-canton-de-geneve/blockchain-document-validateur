
/**
 * Classe qui contient des fonctions pour gérer les tooltip et popover bootstrap.
 * cf. http://getbootstrap.com/javascript/#tooltips et http://getbootstrap.com/javascript/#popovers
 * @author cheutonh
 * @since 27/10/2016
 */

function getHelpHtml(elt, id) {
	var selector = elt;
	// Recherche le prochain élément qui correspond à l'id donné
	var value = selector.nextAll(id).html();
	// Si rien trouvé, remonte d'un niveau en utilisant le parent de l'élément comme sélecteur
	if (value == null) {
		selector = elt.parent();
	}
	value = selector.nextAll(id).html();
	return value;
}
 
(function( $ ) {
 
	/** Fonction qui permet de gérer un tooltip bootstrap */
    $.fn.tooltipify = function( options ) {
 
		// Détermine les paramètres par défaut
		var defaultOptions = {
            trigger: 'focus',
            html: true,
			container: 'body',
			title: function () {
				// Va chercher le titre du tooltip dans le span.help-block suivant
				return getHelpHtml($(this), 'span.help-block');
			},
			// Adapte la position du tooltip suivant la largeur du contenu html
			placement: function(tip, element) {
				var width = $( document ).width();
				if (width <= 953) {
					return "bottom";
				}
				return "right";
			}
        };
		
		// Surchage les paramètres par défaut par le contenu du paramètre "options"
		var settings = $.extend( {}, defaultOptions, options );
 
		// Affiche le tooltip
        this.tooltip({
			trigger: settings.trigger,
			html: settings.html,
			container: settings.container,
			title: settings.title,
			placement: settings.placement
        });
 
        return this;
    };
	
    /** Fonction qui permet de gérer un popover bootstrap */
	$.fn.popoverify = function( options ) {
		
		// Détermine les paramètres par défaut
		var defaultOptions = {
            trigger: 'focus',
            html: true,
			container: 'body',
			title: function () {
				// Va chercher le titre du tooltip dans le span.help-block-title suivant
				return getHelpHtml($(this), 'span.help-block-title');
			},
			content: function () {
				// Va chercher le contenu du tooltip dans le span.help-block-content suivant
				return getHelpHtml($(this), 'span.help-block-content');
			},
			// Adapte la position du tooltip suivant la largeur du contenu html
			placement: function(tip, element) {
				var width = $( document ).width();
				if (width <= 953) {
					return "bottom";
				}
				return "right";
			}
        };
		
		// Surchage les paramètres par défaut par le contenu du paramètre "options"
		var settings = $.extend( {}, defaultOptions, options );
		
		// Affiche le popover
		this.popover({
			trigger: settings.trigger,
			html: settings.html,
			container: settings.container,
			title: settings.title,
			content: settings.content,
			placement: settings.placement
		});
		
		return this;
	};
 
}( jQuery ));
