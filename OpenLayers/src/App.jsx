import { Map, View } from "ol";
import TileLayer from "ol/layer/Tile";
import { createSignal, onMount } from "solid-js";
import OSM from "ol/source/OSM";

import proj4 from "proj4";
import { register } from "ol/proj/proj4";
import { get } from "ol/proj";

import "./app.scss";
import 'ol/ol.css';

import axios from "axios";


proj4.defs("EPSG:2154", "+proj=lcc +lat_0=46.5 +lon_0=3 +lat_1=49 +lat_2=44 +x_0=700000 +y_0=6600000 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m +no_defs +type=crs");
register(proj4)
const projectionL93 = get('EPSG:2154')

function App() {

	var map;

	var mapZone;


	var background = new TileLayer({
		source: new OSM()
	});

	const [getSearch, setSearch] = createSignal("");

	const searchAddress = async () => {
		try {
			let res = await axios.get(`https://api-adresse.data.gouv.fr/search/?q=${getSearch().replace(/ /gi, '+')}`)

			if (res.status === 200){
				console.log(res.data)
			}else throw new Error("Erreur lors de la recherche de l'adresse");
		} catch (e) {
			console.error(e);
		}
	}

	onMount(() => {
		map = new Map({
			view: new View({
				center: [851869.4, 6407969.2],
				zoom: 10,
				projection: projectionL93,
			}),
			layers: [background],
			target: mapZone,
		});
	})


	return (
		<div class="container">

			<input type="search"
				value={getSearch()}
				onInput={(e) => setSearch(e.target.value)}
				placeholder="Entez votre adresse" />

			<button disabled={getSearch() == ''}
				onclick={() => searchAddress()} > 🔍 </button>

			<div ref={mapZone} class="map-zone" />
		</div>
	);

}

export default App;
