Feature: Gestión de tareas con técnica Pomodoro
  Como usuario del sistema Doro
  Quiero gestionar mis tareas usando la técnica Pomodoro
  Para mejorar mi productividad

  Scenario: Crear nuevo usuario
    Given un usuario no registrado con email "usuario@test.com"
    When envío una petición POST a "/users" con los datos del usuario
    Then el sistema debe responder con código 200

  Scenario: Obtener usuario por email
    Given dado un usuario registrado con email "test@example.com"
    When cuando busco el usuario por email
    Then el sistema debe responder con código 200

  Scenario: Obtener listado de backlogs
    Given dado un usuario registrado con email "test@example.com"
    When realizo una petición GET a "/backlogs"
    Then el sistema debe responder con código 200

  Scenario: Obtener listado de tareas
      Given dado un usuario registrado con email "test@example.com"
      When realizo una petición GET a "/tasks"
      Then el sistema debe responder con código 200
