sap.ui.controller("dynamic_textbox.Text", {

	/**
	 * Called when a controller is instantiated and its View controls (if
	 * available) are already created. Can be used to modify the View before it
	 * is displayed, to bind event handlers and do other one-time
	 * initialization.
	 * 
	 * @memberOf dynamic_textbox.Text
	 */
	onInit : function() {
		var oModel = new sap.ui.model.json.JSONModel({
			texts : []
		});
		this.getView().setModel(oModel);
		
		/*var binding = new sap.ui.model.Binding(oModel, "/texts", oModel.getContext("/texts"));
		binding.attachChange(function() {
			var data = oModel.getData("/texts");
			data.push({text : "TEST"});
		});*/
	},
	_onButtonPress : function(evt){
		var oModel = this.getView().getModel();
		var data = oModel.getProperty("/texts");
		data.push({text : "TEST"});
		oModel.refresh();
	}

/**
 * Similar to onAfterRendering, but this hook is invoked before the controller's
 * View is re-rendered (NOT before the first rendering! onInit() is used for
 * that one!).
 * 
 * @memberOf dynamic_textbox.Text
 */
// onBeforeRendering: function() {
//
// },
/**
 * Called when the View has been rendered (so its HTML is part of the document).
 * Post-rendering manipulations of the HTML could be done here. This hook is the
 * same one that SAPUI5 controls get after being rendered.
 * 
 * @memberOf dynamic_textbox.Text
 */
// onAfterRendering: function() {
//
// },
/**
 * Called when the Controller is destroyed. Use this one to free resources and
 * finalize activities.
 * 
 * @memberOf dynamic_textbox.Text
 */
// onExit: function() {
//
// }
});