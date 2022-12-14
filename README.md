# Лабораторная работа 2 по дисциплине "Технологии программирования"

**Изучение фреймворка MVC**

Выполнил: студент группы САПР-1.4 Семёночкин Дмитрий Олегович

## Цели

1. Познакомиться c моделью MVC, ее сущностью и основными фреймворками на ее основе.
2. Разобраться с сущностями «модель», «контроллер», «представление», их функциональным назначением.
3. Получить навыки разработки веб-приложений с использованием MVC-фреймворков, написания модульных тестов к ним;
4. Получить навыки управления автоматизированным тестированием и разворачиванием программного обеспечения, расположенного в системе Git.

## Задачи
1. Выберите для Вашего проекта тип лицензии и добавьте файл с лицензией в проект.
2. Перенести готовый проект по заданиию на язык программирования go
3. Доработайте проект магазина, добавив в него новую функциональность и информацию в базу данных в соответствии с типом магазина (согласно индивидуальному варианту, см. таблицу). Составьте модульные тесты к проекту, постарайтесь покрыть тестами максимально возможный объем кода. Для работы с этим заданием создайте новую ветку кода на основе главной и фиксируйте в нее весь программный код в процессе разработки. Добейтесь выполнения всех тестов проекта, после чего объедините текущую ветку кода с главной.
4. Проанализируйте полученные результаты и сделайте выводы.

## Постановка задачи
Дан проект сайта интернет-магазина добавить новую функциональность, в соответствии с индивидуальным вариантом:

- сформировать ассортимент магазина электроники 
- реализовать:
    - учёт количества товара
    - добавить скидку на товар в 15% после покупки 10 единиц товара
    - реализовать скидку на определенный период времени

## Индивидуальный вариант

| Вариант | Тип магазина          | Функциональность приложения                                                                                                                                            |
|---------|-----------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 8       | Магазин электроники   | В магазине имеется определенное количество товара каждого вида. После покупки 10 экземпляров любого товара (возможно, разными покупателями) его цена возрастает на 15%.|

## Используемые язык и библиотеки

Используется язык `Go` с фреймворком `gin`, ORM `gorm` и библиотека для тестирования testify 

## Краткое описание проекта

Проект представляет собой интернет магазин. По индивидуальному в. Количестов товара ограничено и уменьшается при покупке. Если товар купили 10 раз, его цена уменьшается на 15%.

## Выводы

Выполнив лабораторную работу я реализовал приложение в модели MVC с использованием фреймворка `gin` на языке `go`
