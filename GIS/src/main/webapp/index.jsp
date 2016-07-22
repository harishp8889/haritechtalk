<!DOCTYPE html>
<html>
<head>
<meta charset="ISO-8859-1">
<title>Insert title here</title>
<link rel="stylesheet" href="http://localhost:9090/API/OL3/css/ol.css"
	type="text/css">
<script src="http://localhost:9090/API/OL3/build/ol.js"></script>
<script src="http://localhost:9090/API/JQuery/core/jquery-2.2.3.min.js"></script>
</head>
<body>
	<div id='map'></div>
	<form class="form-inline">
		<label>Geometry type &nbsp;</label> <select id="type">
			<option value="Point">Point</option>
			<option value="LineString">LineString</option>
			<option value="Polygon">Polygon</option>
		</select>
		<button value="Save" title="Save" onclick="save()">Save</button>
	</form>
	<script type="text/javascript">
		var raster = new ol.layer.Tile({
			source : new ol.source.MapQuest({
				layer : 'osm'
			})
		});
		var wmsSource = new ol.source.ImageWMS({
			url : 'http://localhost:9090/geoserver/wfs',
			params : {
				'LAYERS' : 'haritechtalk:Employee',
				'LAYERS' : 'haritechtalk:Flight'
			},
			serverType : 'geoserver',
			crossOrigin : 'anonymous'
		});

		var wmsLayer = new ol.layer.Image({
			source : wmsSource
		});
		var map = new ol.Map({
			layers : [ raster, wmsLayer ],
			target : 'map',
			view : new ol.View({
				center : [ -11000000, 4600000 ],
				zoom : 4
			})
		});

		var features = new ol.Collection();
		var featureOverlay = new ol.layer.Vector({
			source : new ol.source.Vector({
				features : features
			}),
			style : new ol.style.Style({
				fill : new ol.style.Fill({
					color : 'rgba(255, 255, 255, 0.2)'
				}),
				stroke : new ol.style.Stroke({
					color : '#ffcc33',
					width : 2
				}),
				image : new ol.style.Circle({
					radius : 7,
					fill : new ol.style.Fill({
						color : '#ffcc33'
					})
				})
			})
		});
		featureOverlay.setMap(map);

		var modify = new ol.interaction.Modify({
			features : features,
			// the SHIFT key must be pressed to delete vertices, so
			// that new vertices can be drawn at the same position
			// of existing vertices
			deleteCondition : function(event) {
				return ol.events.condition.shiftKeyOnly(event)
						&& ol.events.condition.singleClick(event);
			}
		});
		map.addInteraction(modify);

		var draw; // global so we can remove it later
		var typeSelect = document.getElementById('type');

		function addInteraction() {
			draw = new ol.interaction.Draw({
				features : features,
				type : /** @type {ol.geom.GeometryType} */
				(typeSelect.value)
			});
			draw.addEventListener("drawend", function(evt) {
				format = new ol.format["GeoJSON"]();
				var data;
				try {
					// convert the data of the vector_layer into the chosen format
					data = format.writeFeatures(features);
				} catch (e) {
					// at time of creation there is an error in the GPX format (18.7.2014)
					console.log(e.name + ": " + e.message);
					return;
				}
				// format is JSON
				console.log(JSON.stringify(data, null, 4));
			});
			map.addInteraction(draw);
		}

		/**
		 * Handle change event.
		 */
		typeSelect.onchange = function() {
			map.removeInteraction(draw);
			addInteraction();
		};

		addInteraction();
	</script>
</body>
</html>