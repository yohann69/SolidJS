import { For, Show, createEffect, createMemo, createSignal, onMount } from 'solid-js';
import axios from 'axios';


export default function(props){
    const [getTodo, setTodo] = createSignal("")

	const [getTodos, setTodos] = createSignal([])

    onMount(async () => {
        console.log(`Démarage du composant ${props.name}`)
        try{
            const response = await axios.get('https://swapi.dev/api/planets')
            if(response.status === 200){
                setTodos(response.data.results.map(e => e.name))
            } else throw new Error()

        }catch(e){
            console.error(e)
        }
    })

/* 	createEffect(() => {
		if (getTodo() !== "") {
			console.log("Il y a eu du changement");
			console.log(getTodo())
		}
	}) */

/* 	const getNewTodo = createMemo(() => {
		return getTodo().split('').reverse().join('')
	}) */


	return (
		<>
            <h1>TodoList de {props.name}</h1>
			<input
				type="text"
				value={getTodo()}
				onInput={
					(e) => setTodo(e.target.value)
				}
			/>
			<button disabled={getTodo() === ""} onclick={() => {
				setTodos([...getTodos(), getTodo()])
				setTodo("")
			}}>Ajouter</button>
{/* 			<span>{getNewTodo()}</span> */}

			<Show when={getTodos().length > 0} fallback={<span>La liste est vide</span>}>
				<ul>Choses à faire:
					<For each={getTodos()}>{(todo, index) =>
						<li>{index()} : {todo} - <button onclick={() => setTodos(oldTodos => oldTodos.filter((v, i) => i != index()))}>X</button></li>}
					</For>
				</ul>
			</Show>
		</>
	);
}