1. Напишите функцию, которая будет получать на вход строку с поисковым запросом (string) и массив ссылок на страницы, по которым стоит произвести поиск ([]string). Результатом работы функции должен быть массив строк со ссылками на страницы, на которых обнаружен поисковый запрос. Функция должна искать точное соответствие фразе в тексте ответа от сервера по каждой из ссылок.
   Подсказка: это задача на поиск последовательности в массиве.
 
     findOnSite.go search.go search2.go выполняют эту задачу.  
     
     go run ./findOnSite -query="строка для поиска" -query="https://1.com,https://2.com,..."
     
 2. * Напишите функцию, которая получает на вход публичную ссылку на файл с «Яндекс.Диска» и сохраняет полученный файл на диск пользователя.
 yaDisk.go
 
 GET https://cloud-api.yandex.net/v1/disk/resources/download?path=%2FДокумент.docx
 где аргумент для path - путь к файлу на диске.
 в авторизации нужно указать OAuth