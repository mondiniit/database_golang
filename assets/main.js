const btn_send = document.querySelector('#btn_send')
const spots = document.querySelectorAll('.inputs.request')

spots[0].focus()
this.selectionStart = this.selectionEnd

btn_send.addEventListener('click', (e) => {
	e.preventDefault()
	enviar()
})

spots.forEach((input) => input.addEventListener('keyup', (e) => {
	if (input.value.length > 0 && input.nextElementSibling) {
		input.nextElementSibling.focus()
		return
	}
}))
spots.forEach((input) => input.addEventListener('keyup', (e) => {
	if (e.keyCode == 8 && input.value.length == 0) {
		if (input.previousElementSibling) {
			input.previousElementSibling.focus()
		}
		return
	}
	if (e.keyCode == 37 && input.previousElementSibling) {
		input.previousElementSibling.focus()
	}
	if (e.keyCode == 39 && input.nextElementSibling) {
		input.nextElementSibling.focus()
	}
}))

spots.forEach((input) => input.addEventListener('keypress', (e) => {
	if (e.keyCode != 13 && e.keyCode != 9) {
		input.value = input.value.replace(/./g, '')
	}
}))

async function enviar() {
	const response = []
	const request_list = document.querySelectorAll('.inputs.request')
	request_list.forEach((request) => response.push(request.value))
	let tips = await fecth_expression(response)
	fill_igual(response, tips)
	fill_tips(tips)
	clear_inputs(request_list)
	request_list[0].focus()
}

async function fecth_expression(req) {
	const update = {
		input: req.toString(),
	};
	const options = {
		method: 'POST',
		headers: {
		'Content-Type': 'application/json',
		},
		body: JSON.stringify(update),
	};
	let response = await fetch(`http://localhost:8080`, options)
	.then(encoded => encoded.json())
	.then(response => response)
	return(response)
}

function fill_tips(tips) {
	const element_tips = document.querySelectorAll(".inputs.tips")
	const inputs = document.querySelectorAll(".inputs.request")
	const element = document.querySelector('.d_msg')
	element_tips.forEach((tip, index) => tip.classList.remove("winner"))
	if (tips.content == 'CCCCCC') {
		element_tips.forEach((tip, index) => tip.classList.add("winner"))
		inputs.forEach((input, index) => input.disabled = true)
	}
	if (!tips.error) {
		element_tips.forEach((tip, index) => tip.innerHTML = tips.content[index])
		element.innerHTML = ""
	} else {
		element.innerHTML = tips.error
		element.classList.add("warning")
	}
}

function fill_igual(req, tips) {
	const element_tips = document.querySelectorAll(".inputs.last_input")
	if (!tips.error) {
		element_tips.forEach((tip, index) => tip.innerHTML = req[index])
	}
}

function clear_inputs(lst) {
	const inputs = document.querySelectorAll(".inputs.request")
	if (inputs[0].disabled == true) {
		inputs.forEach((input) => input.classList.add("winner"))
	} else {
		inputs.forEach((input, index) => input.value = lst[index].value)
	}
}
