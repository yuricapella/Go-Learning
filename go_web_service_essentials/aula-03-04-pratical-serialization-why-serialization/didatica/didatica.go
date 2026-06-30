package didatica

import "fmt"

func DemonstrarSerializacao() {
	fmt.Println("=== AULAS 03-04: Serialization ===")
	fmt.Println()

	fmt.Println("--- O que e serializacao? ---")
	fmt.Println("Serializar e transformar um valor da linguagem em bytes.")
	fmt.Println("Desserializar e fazer o caminho contrario: ler bytes e reconstruir um valor.")
	fmt.Println()

	fmt.Println("--- Por que isso existe? ---")
	fmt.Println("Memoria do programa nao atravessa rede, arquivo ou HTTP do jeito que esta.")
	fmt.Println("Para enviar um int, float, string ou struct, voce precisa combinar um formato de bytes que outro processo consiga entender.")
	fmt.Println()

	Exemplo1BytesEEndianess()
	Exemplo2JSONEmHandlerHTTP()
	Exemplo3FluxoDeserializeWorkSerialize()
	PontosImportantes()
}

func Exemplo1BytesEEndianess() {
	fmt.Println("--- Exemplo 1: bytes e endianess ---")
	fmt.Println("Numeros precisam virar bytes em alguma ordem.")
	fmt.Println("Big endian guarda o byte mais significativo primeiro.")
	fmt.Println("Little endian guarda o byte menos significativo primeiro.")
	fmt.Println("O importante nao e decorar o nome: e saber que quem escreve e quem le precisam concordar na ordem.")
	fmt.Println()
}

func Exemplo2JSONEmHandlerHTTP() {
	fmt.Println("--- Exemplo 2: JSON em HTTP ---")
	fmt.Println("Em APIs HTTP, JSON costuma ser o formato de serializacao.")
	fmt.Println("O cliente envia JSON no body; o servidor faz decode para uma struct.")
	fmt.Println("Depois de trabalhar, o servidor faz encode de uma resposta para JSON.")
	fmt.Println()
}

func Exemplo3FluxoDeserializeWorkSerialize() {
	fmt.Println("--- Exemplo 3: fluxo pratico ---")
	fmt.Println("1. Deserialize: `json.NewDecoder(r.Body).Decode(&valor)`.")
	fmt.Println("2. Work: execute a regra da aplicacao.")
	fmt.Println("3. Serialize: `json.NewEncoder(w).Encode(resposta)`.")
	fmt.Println()
}

func PontosImportantes() {
	fmt.Println("--- Pontos importantes ---")
	fmt.Println("- Bytes so fazem sentido quando existe um formato combinado.")
	fmt.Println("- Use `encoding/json` para JSON em vez de montar string manualmente.")
	fmt.Println("- Feche o body da request quando terminar de ler.")
	fmt.Println("- Defina `Content-Type: application/json` em respostas JSON.")
}
