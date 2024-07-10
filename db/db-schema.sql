CREATE TABLE Candidates (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    telefono VARCHAR(15),
    direccion VARCHAR(255),
    fecha_nacimiento DATE,
    estado VARCHAR(50)
);

CREATE TABLE JobPositions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo_puesto VARCHAR(100) NOT NULL,
    descripcion TEXT,
    requisitos TEXT,
    fecha_publicacion DATE,
    fecha_cierre DATE,
    estado VARCHAR(50)
);

CREATE TABLE Applications (
    id INT AUTO_INCREMENT PRIMARY KEY,
    candidato_id INT,
    puesto_id INT,
    fecha_solicitud DATE,
    estado_solicitud VARCHAR(50),
    FOREIGN KEY (candidato_id) REFERENCES Candidates(id),
    FOREIGN KEY (puesto_id) REFERENCES JobPositions(id)
);



CREATE TABLE WorkExperience (
    id INT AUTO_INCREMENT PRIMARY KEY,
    candidato_id INT,
    empresa VARCHAR(100),
    cargo VARCHAR(100),
    fecha_inicio DATE,
    fecha_fin DATE,
    descripcion TEXT,
    FOREIGN KEY (candidato_id) REFERENCES Candidates(id)
);

CREATE TABLE Education (
    id INT AUTO_INCREMENT PRIMARY KEY,
    candidato_id INT,
    institucion VARCHAR(100),
    titulo VARCHAR(100),
    fecha_inicio DATE,
    fecha_fin DATE,
    FOREIGN KEY (candidato_id) REFERENCES Candidates(id)
);

CREATE TABLE Interviewers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    departamento VARCHAR(100),
    titulo_puesto VARCHAR(100)
);

CREATE TABLE Interviews (
    id INT AUTO_INCREMENT PRIMARY KEY,
    candidato_id INT,
    puesto_id INT,
    fecha_hora DATETIME,
    entrevistador_id INT,
    notas TEXT,
    resultado VARCHAR(50),
    FOREIGN KEY (candidato_id) REFERENCES Candidates(id),
    FOREIGN KEY (puesto_id) REFERENCES JobPositions(id),
    FOREIGN KEY (entrevistador_id) REFERENCES Interviewers(id)
);

CREATE TABLE Skills (
    id INT AUTO_INCREMENT PRIMARY KEY,
    candidato_id INT,
    habilidad VARCHAR(100),
    nivel VARCHAR(50),
    FOREIGN KEY (candidato_id) REFERENCES Candidates(id)
);



CREATE TABLE CandidateJobApplications (
    id INT AUTO_INCREMENT PRIMARY KEY,
    candidato_id INT,
    puesto_id INT,
    fecha_solicitud DATE,
    estado_solicitud VARCHAR(50),
    FOREIGN KEY (candidato_id) REFERENCES Candidates(id),
    FOREIGN KEY (puesto_id) REFERENCES JobPositions(id)
);

CREATE TABLE Resumes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    candidato_id INT,
    nombre_archivo VARCHAR(255),
    tipo_archivo VARCHAR(50),
    url_archivo VARCHAR(255),
    fecha_subida DATE,
    FOREIGN KEY (candidato_id) REFERENCES Candidates(id)
);


