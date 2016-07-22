sap.ui.jsview("root.main", {

	/**
	 * Specifies the Controller belonging to this View. In the case that it is
	 * not implemented, or that "null" is returned, this View does not have a
	 * Controller.
	 * 
	 * @memberOf modules.when.controller.displayWhen
	 */
	getControllerName : function() {
		return "root.main";
	},

	/**
	 * Is initially called once after the Controller has been instantiated. It
	 * is the place where the UI is constructed. Since the Controller is given
	 * to this method, its event handlers can be attached right away.
	 * 
	 * @memberOf modules.when.controller.displayWhen
	 */
	createContent : function(oController) {
		var input = new sap.m.Input({
			value : "Test"
		});
		return input;
	},

});