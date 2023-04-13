import { Map, View } from "ol";
import TileLayer from "ol/layer/Tile";
import { onMount } from "solid-js";
import OSM from "ol/source/OSM";

import "./app.scss";

function App() {

	var map;

	var mapZone;

	var background = new TileLayer({
		source: new OSM()
	});

	onMount(() => {
		map = new Map({
			view: new View({
				center: [0, 0],
				zoom: 1
			}),
			layers: [background],
			target: mapZone,
		});
	})


	return (
		<div class="container">
			<div ref={mapZone} class="map-zone"/>
		</div>
	);

}

export default App;
