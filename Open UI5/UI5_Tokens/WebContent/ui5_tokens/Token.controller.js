sap.ui.controller("ui5_tokens.Token", {
	suggestModel : null,
	valsArr : [],
	dict : null,
	/**
	 * Called when a controller is instantiated and its View controls (if
	 * available) are already created. Can be used to modify the View before it
	 * is displayed, to bind event handlers and do other one-time
	 * initialization.
	 * 
	 * @memberOf ui5_tokens.Token
	 */
	onInit : function() {
		var oMultiInput1 = this.getView().byId("m1");
		// *** add text change validator
		oMultiInput1.addValidator(function(args) {
			return new sap.m.Token({
				key : args.text,
				text : args.text
			});
		});
		var oData = {
			names : [ {
				name : "ABC"
			}, {
				name : "XYZ"
			}, {
				name : "Test"
			} ]
		};
		this.suggestModel = new sap.ui.model.json.JSONModel(oData);

	},
	_enableSuggest : function(oEvent) {
		if (oEvent.getParameter("suggestValue").length > 0) {
			oEvent.getSource().setModel(this.suggestModel);
			oEvent.getSource().bindAggregation("suggestionItems", "/names",
					new sap.ui.core.Item({
						text : "{name}"
					}));
		}
	},
	_multiInpChanged : function(oEvent) {
		console.log(oEvent);
		switch (oEvent.getParameter("type")) {
		case "added":
			var keys = oEvent.getParameter("token").getKey();
			var val = oEvent.getParameter("token").getKey();
			this.dict = {
				keys : val
			};
			this.valsArr.push(this.dict);
			break;
		case "removed":
			var key = oEvent.getParameter("token").getKey();
			delete this.dict[key];
			break;
		}
		// lookup:
		/*
		 * var second = dict["2"]; // update: dict["2"] = false; // add:
		 * dict["4"] = true; // delete: delete dict["2"];
		 */
	}

/**
 * Similar to onAfterRendering, but this hook is invoked before the controller's
 * View is re-rendered (NOT before the first rendering! onInit() is used for
 * that one!).
 * 
 * @memberOf ui5_tokens.Token
 */
// onBeforeRendering: function() {
//
// },
/**
 * Called when the View has been rendered (so its HTML is part of the document).
 * Post-rendering manipulations of the HTML could be done here. This hook is the
 * same one that SAPUI5 controls get after being rendered.
 * 
 * @memberOf ui5_tokens.Token
 */
// onAfterRendering: function() {
//
// },
/**
 * Called when the Controller is destroyed. Use this one to free resources and
 * finalize activities.
 * 
 * @memberOf ui5_tokens.Token
 */
// onExit: function() {
//
// }
});