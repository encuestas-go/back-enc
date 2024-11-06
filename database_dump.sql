-- MySQL dump 10.13  Distrib 8.3.0, for macos14.2 (arm64)
--
-- Host: 127.0.0.1    Database: ENCUESTA
-- ------------------------------------------------------
-- Server version	8.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `ENCUESTA_ACTIVIDAD`
--

DROP TABLE IF EXISTS `ENCUESTA_ACTIVIDAD`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ENCUESTA_ACTIVIDAD` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ID_USUARIO` int DEFAULT NULL,
  `JUEGOS_PREFERIDOS` varchar(100) DEFAULT NULL,
  `PASATIEMPOS` varchar(100) DEFAULT NULL,
  `DEPORTE_INTERES` varchar(100) DEFAULT NULL,
  `FRECUENCIA_EJERCICIO` varchar(100) DEFAULT NULL,
  `TIPO_TALLERES` varchar(100) DEFAULT NULL,
  `EVENTOS_SOCIALES` varchar(100) DEFAULT NULL,
  `FECHA` date DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_USUARIO` (`ID_USUARIO`),
  CONSTRAINT `ENCUESTA_ACTIVIDAD_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ENCUESTA_ACTIVIDAD`
--

LOCK TABLES `ENCUESTA_ACTIVIDAD` WRITE;
/*!40000 ALTER TABLE `ENCUESTA_ACTIVIDAD` DISABLE KEYS */;
INSERT INTO `ENCUESTA_ACTIVIDAD` (`ID`, `ID_USUARIO`, `JUEGOS_PREFERIDOS`, `PASATIEMPOS`, `DEPORTE_INTERES`, `FRECUENCIA_EJERCICIO`, `TIPO_TALLERES`, `EVENTOS_SOCIALES`, `FECHA`) VALUES (1,5,'Online','Baile,Tocar algún instrumento,Pintar,Hacer ejercicio','Volleyball','Mensual','Habilidades blandas,Bienestar y salud','Conciertos,Bailes,Charlas/Conferencias','2024-06-06'),(2,16,'Videojuegos','Leer,Salir a caminar','Atletismo','No puedo','Habilidades blandas,Desarrollo personal','Literatura/poesía,Charlas/Conferencias','2024-06-06'),(3,18,'Online','Baile,Tocar algún instrumento,Pintar,Salir a caminar','Basketball','Raramente','Emprendimiento,Desarrollo personal','Festivales,Exposiciones de arte,Charlas/Conferencias','2024-06-06'),(4,19,'Online','Pintar,Salir a caminar,Series o películas','Volleyball','Raramente','Habilidades blandas,Técnicos(Con relación a la carrera)','Exposiciones de arte,Charlas/Conferencias','2024-06-06'),(5,20,'Online','Tocar algún instrumento,Dibujar,Salir a caminar','Ciclismo','Raramente','Habilidades blandas,Desarrollo personal','Conciertos,Exposiciones de arte,Charlas/Conferencias','2024-06-06');
/*!40000 ALTER TABLE `ENCUESTA_ACTIVIDAD` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
/*!50032 DROP TRIGGER IF EXISTS AGREGAR_FECHA_ACTUAL */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `AGREGAR_FECHA_ACTUAL` BEFORE INSERT ON `ENCUESTA_ACTIVIDAD` FOR EACH ROW BEGIN
    SET NEW.FECHA = CURDATE();
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
/*!50032 DROP TRIGGER IF EXISTS ACTUALIZAR_FECHA_ACTUAL */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `ACTUALIZAR_FECHA_ACTUAL` BEFORE UPDATE ON `ENCUESTA_ACTIVIDAD` FOR EACH ROW BEGIN
    SET NEW.FECHA = CURDATE();
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `ENCUESTA_INFRAESTRUCTURA_HOGAR`
--

DROP TABLE IF EXISTS `ENCUESTA_INFRAESTRUCTURA_HOGAR`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ENCUESTA_INFRAESTRUCTURA_HOGAR` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ID_USUARIO` int DEFAULT NULL,
  `ZONA` varchar(100) DEFAULT NULL,
  `PERMANENCIA` varchar(100) DEFAULT NULL,
  `ESTADO_INFRAESTRUCTURA` varchar(100) DEFAULT NULL,
  `TIPO_SUELO` varchar(100) DEFAULT NULL,
  `TIPO_TECHO` varchar(100) DEFAULT NULL,
  `TIPO_PARED` varchar(100) DEFAULT NULL,
  `NUMERO_INTEGRANTES` int DEFAULT NULL,
  `NUMERO_HABITACIONES` int DEFAULT NULL,
  `EQUIPAMIENTO_HOGAR` varchar(100) DEFAULT NULL,
  `SERVICIOS_BASICOS` varchar(100) DEFAULT NULL,
  `OTRAS_PROPIEDADES` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ENCUESTA_INFRAESTRUCTURA_HOGAR_ibfk_1` (`ID_USUARIO`),
  CONSTRAINT `ENCUESTA_INFRAESTRUCTURA_HOGAR_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ENCUESTA_INFRAESTRUCTURA_HOGAR`
--

LOCK TABLES `ENCUESTA_INFRAESTRUCTURA_HOGAR` WRITE;
/*!40000 ALTER TABLE `ENCUESTA_INFRAESTRUCTURA_HOGAR` DISABLE KEYS */;
INSERT INTO `ENCUESTA_INFRAESTRUCTURA_HOGAR` (`ID`, `ID_USUARIO`, `ZONA`, `PERMANENCIA`, `ESTADO_INFRAESTRUCTURA`, `TIPO_SUELO`, `TIPO_TECHO`, `TIPO_PARED`, `NUMERO_INTEGRANTES`, `NUMERO_HABITACIONES`, `EQUIPAMIENTO_HOGAR`, `SERVICIOS_BASICOS`, `OTRAS_PROPIEDADES`) VALUES (18,16,'Urbano','Largo Plazo','Necesita Reparaciones','Mármol','Ladrillo','Piedra',11,2,'Televisión','Agua',1),(19,18,'Urbano','Mediano Plazo','Regular','Alfombra','Piedra','Piedra',33,3,'Refrigerador,Lavadora','Electricidad,Calefacción',1),(20,19,'Urbano','Mediano Plazo','Regular','Alfombra','Piedra','Piedra',3,3,'Televisión,Refrigerador','Electricidad',1),(21,20,'Urbano','Mediano Plazo','Necesita Reparaciones','Alfombra','Piedra','Madera',2,2,'Televisión,Ventilador','Electricidad,Agua',1);
/*!40000 ALTER TABLE `ENCUESTA_INFRAESTRUCTURA_HOGAR` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ENCUESTA_NIVEL_DEMOGRAFICO`
--

DROP TABLE IF EXISTS `ENCUESTA_NIVEL_DEMOGRAFICO`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ENCUESTA_NIVEL_DEMOGRAFICO` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ID_USUARIO` int DEFAULT NULL,
  `TIPO_VIVIENDA` varchar(100) DEFAULT NULL,
  `TIPO_CONDICION` varchar(100) DEFAULT NULL,
  `TRANSPORTE_PROPIO` tinyint(1) DEFAULT NULL,
  `MONTO_INGRESOS` float DEFAULT NULL,
  `NUM_INTEGRANTES_TRABAJAN` int DEFAULT NULL,
  `NUM_INTEGRANTES_MENOR_EDAD` int DEFAULT NULL,
  `DESPENSA_MENSUAL` float DEFAULT NULL,
  `APOYOS_GOBIERNO` tinyint(1) DEFAULT NULL,
  `FECHA` date DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_USUARIO` (`ID_USUARIO`),
  CONSTRAINT `ENCUESTA_NIVEL_DEMOGRAFICO_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ENCUESTA_NIVEL_DEMOGRAFICO`
--

LOCK TABLES `ENCUESTA_NIVEL_DEMOGRAFICO` WRITE;
/*!40000 ALTER TABLE `ENCUESTA_NIVEL_DEMOGRAFICO` DISABLE KEYS */;
INSERT INTO `ENCUESTA_NIVEL_DEMOGRAFICO` (`ID`, `ID_USUARIO`, `TIPO_VIVIENDA`, `TIPO_CONDICION`, `TRANSPORTE_PROPIO`, `MONTO_INGRESOS`, `NUM_INTEGRANTES_TRABAJAN`, `NUM_INTEGRANTES_MENOR_EDAD`, `DESPENSA_MENSUAL`, `APOYOS_GOBIERNO`, `FECHA`) VALUES (1,5,'Cuarto','Propia',0,9500,3,3,1222,1,'2024-06-06'),(2,16,'Departamento','Prestado',1,30000,3,2,2222,1,'2024-06-06'),(3,18,'Cuarto','Prestado',1,9500,3,3,333,1,'2024-06-06'),(4,19,'Vecindad','Prestado',1,30000,2,2,3333,0,'2024-06-06'),(5,20,'Cuarto','Propia',1,20000,1,2,3333,1,'2024-06-06');
/*!40000 ALTER TABLE `ENCUESTA_NIVEL_DEMOGRAFICO` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
/*!50032 DROP TRIGGER IF EXISTS AGREGAR_FECHA_ACTUAL_DEMOGRAFICO */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `AGREGAR_FECHA_ACTUAL_DEMOGRAFICO` BEFORE INSERT ON `ENCUESTA_NIVEL_DEMOGRAFICO` FOR EACH ROW BEGIN
    SET NEW.FECHA = CURDATE();
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
/*!50032 DROP TRIGGER IF EXISTS ACTUALIZAR_FECHA_ACTUAL_DEMOGRAFICO */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `ACTUALIZAR_FECHA_ACTUAL_DEMOGRAFICO` BEFORE UPDATE ON `ENCUESTA_NIVEL_DEMOGRAFICO` FOR EACH ROW BEGIN
    SET NEW.FECHA = CURDATE();
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `ENCUESTA_NIVEL_ECONOMICO`
--

DROP TABLE IF EXISTS `ENCUESTA_NIVEL_ECONOMICO`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ENCUESTA_NIVEL_ECONOMICO` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ID_USUARIO` int DEFAULT NULL,
  `SITUACION_ACTUAL` varchar(100) DEFAULT NULL,
  `NOMBRE_EMPLEO` varchar(100) DEFAULT NULL,
  `EMPRESA_ESTABLECIMIENTO` varchar(100) DEFAULT NULL,
  `TIPO_EMPLEO` varchar(100) DEFAULT NULL,
  `SALARIO` float DEFAULT NULL,
  `TIPO_MONTO` varchar(100) DEFAULT NULL,
  `TIPO_PRESTACIONES` varchar(100) DEFAULT NULL,
  `FECHA` date DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_USUARIO` (`ID_USUARIO`),
  CONSTRAINT `ENCUESTA_NIVEL_ECONOMICO_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ENCUESTA_NIVEL_ECONOMICO`
--

LOCK TABLES `ENCUESTA_NIVEL_ECONOMICO` WRITE;
/*!40000 ALTER TABLE `ENCUESTA_NIVEL_ECONOMICO` DISABLE KEYS */;
INSERT INTO `ENCUESTA_NIVEL_ECONOMICO` (`ID`, `ID_USUARIO`, `SITUACION_ACTUAL`, `NOMBRE_EMPLEO`, `EMPRESA_ESTABLECIMIENTO`, `TIPO_EMPLEO`, `SALARIO`, `TIPO_MONTO`, `TIPO_PRESTACIONES`, `FECHA`) VALUES (1,5,'Empleado','Programadores','Coca Cola','Contrato temporal',20000,'Quincenal','Segurado Medico Particular','2024-06-06'),(2,16,'Estudiante','Programadores','Coca Cola','Contrato temporal',20000,'Quincenal','Segurado Medico Particular','2024-06-06');
/*!40000 ALTER TABLE `ENCUESTA_NIVEL_ECONOMICO` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
/*!50032 DROP TRIGGER IF EXISTS AGREGAR_FECHA_ACTUAL_ECONOMICO */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `AGREGAR_FECHA_ACTUAL_ECONOMICO` BEFORE INSERT ON `ENCUESTA_NIVEL_ECONOMICO` FOR EACH ROW BEGIN
    SET NEW.FECHA = CURDATE();
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
/*!50032 DROP TRIGGER IF EXISTS ACTUALIZAR_FECHA_ACTUAL_ECONOMICO */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `ACTUALIZAR_FECHA_ACTUAL_ECONOMICO` BEFORE UPDATE ON `ENCUESTA_NIVEL_ECONOMICO` FOR EACH ROW BEGIN
    SET NEW.FECHA = CURDATE();
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `ENCUESTA_NIVEL_SOCIOECONOMICO`
--

DROP TABLE IF EXISTS `ENCUESTA_NIVEL_SOCIOECONOMICO`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ENCUESTA_NIVEL_SOCIOECONOMICO` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ID_USUARIO` int DEFAULT NULL,
  `NOMBRE_COMPLETO` varchar(100) DEFAULT NULL,
  `FECHA_NACIMIENTO` date DEFAULT NULL,
  `NACIONALIDAD` varchar(100) DEFAULT NULL,
  `SEXO` varchar(100) DEFAULT NULL,
  `EDAD` int DEFAULT NULL,
  `ESTADO_CIVIL` varchar(100) DEFAULT NULL,
  `DIRECCION_RESIDENCIA` varchar(100) DEFAULT NULL,
  `CIUDAD_RESIDENCIA` varchar(100) DEFAULT NULL,
  `CODIGO_POSTAL` int DEFAULT NULL,
  `ENTIDAD_FEDERATIVA` varchar(100) DEFAULT NULL,
  `ESTATUS_SOCIOECONOMICO` varchar(100) DEFAULT NULL,
  `IDIOMA` varchar(100) DEFAULT NULL,
  `GRADO_ESTUDIOS_ASPIRAR` varchar(100) DEFAULT NULL,
  `ULTIMO_GRADO_PADRE` varchar(100) DEFAULT NULL,
  `ULTIMO_GRADO_MADRE` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_USUARIO` (`ID_USUARIO`),
  CONSTRAINT `ENCUESTA_NIVEL_SOCIOECONOMICO_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ENCUESTA_NIVEL_SOCIOECONOMICO`
--

LOCK TABLES `ENCUESTA_NIVEL_SOCIOECONOMICO` WRITE;
/*!40000 ALTER TABLE `ENCUESTA_NIVEL_SOCIOECONOMICO` DISABLE KEYS */;
INSERT INTO `ENCUESTA_NIVEL_SOCIOECONOMICO` (`ID`, `ID_USUARIO`, `NOMBRE_COMPLETO`, `FECHA_NACIMIENTO`, `NACIONALIDAD`, `SEXO`, `EDAD`, `ESTADO_CIVIL`, `DIRECCION_RESIDENCIA`, `CIUDAD_RESIDENCIA`, `CODIGO_POSTAL`, `ENTIDAD_FEDERATIVA`, `ESTATUS_SOCIOECONOMICO`, `IDIOMA`, `GRADO_ESTUDIOS_ASPIRAR`, `ULTIMO_GRADO_PADRE`, `ULTIMO_GRADO_MADRE`) VALUES (22,5,'Christian','2001-12-13','Mexicano','Femenino',22,'Soltero','Lagunilla','Cuernavaca',62037,'Cuernavaca','Media-Alta','Ingles','Maestría','Primaria','Primaria'),(23,16,'AXL','2001-12-13','Mexicano','Femenino',21,'Soltero','Lagunilla','Cuernavaca',62037,'Cuernavaca','Media','Ingles','Maestría','Secundaria','Preparatoria');
/*!40000 ALTER TABLE `ENCUESTA_NIVEL_SOCIOECONOMICO` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ENCUESTA_SATISFACCION`
--

DROP TABLE IF EXISTS `ENCUESTA_SATISFACCION`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ENCUESTA_SATISFACCION` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ID_USUARIO` int DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_USUARIO` (`ID_USUARIO`),
  CONSTRAINT `ENCUESTA_SATISFACCION_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ENCUESTA_SATISFACCION`
--

LOCK TABLES `ENCUESTA_SATISFACCION` WRITE;
/*!40000 ALTER TABLE `ENCUESTA_SATISFACCION` DISABLE KEYS */;
/*!40000 ALTER TABLE `ENCUESTA_SATISFACCION` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ENCUESTA_SERVICIO`
--

DROP TABLE IF EXISTS `ENCUESTA_SERVICIO`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ENCUESTA_SERVICIO` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ID_USUARIO` int DEFAULT NULL,
  `PROVEEDOR_LUZ` tinyint(1) DEFAULT NULL,
  `PROVEEDOR_AGUA` tinyint(1) DEFAULT NULL,
  `PROVEEDOR_INTERNET` varchar(100) DEFAULT NULL,
  `PROVEEDOR_TELEFONO` tinyint(1) DEFAULT NULL,
  `PROVEEDOR_TELEVISION` tinyint(1) DEFAULT NULL,
  `VENCIMIENTO_PAGOS` varchar(100) DEFAULT NULL,
  `PAGOS_ADICIONALES` varchar(100) DEFAULT NULL,
  `GASTOS_SERVICIOS` float DEFAULT NULL,
  `FECHA` date DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_USUARIO` (`ID_USUARIO`),
  CONSTRAINT `ENCUESTA_SERVICIO_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ENCUESTA_SERVICIO`
--

LOCK TABLES `ENCUESTA_SERVICIO` WRITE;
/*!40000 ALTER TABLE `ENCUESTA_SERVICIO` DISABLE KEYS */;
INSERT INTO `ENCUESTA_SERVICIO` (`ID`, `ID_USUARIO`, `PROVEEDOR_LUZ`, `PROVEEDOR_AGUA`, `PROVEEDOR_INTERNET`, `PROVEEDOR_TELEFONO`, `PROVEEDOR_TELEVISION`, `VENCIMIENTO_PAGOS`, `PAGOS_ADICIONALES`, `GASTOS_SERVICIOS`, `FECHA`) VALUES (1,5,1,1,'Totalplay',1,1,'Bimestral','Saldo/Plan Telefónico,Entretenimiento (Netflix|Youtube Premium|Amazon Prime|Spotify...)',1222,'2024-06-06'),(2,16,1,1,'IZZI',1,1,'Cuatrimestral','Saldo/Plan Telefónico,Entretenimiento (Netflix|Youtube Premium|Amazon Prime|Spotify...),Otros',5555,'2024-06-06'),(3,18,1,1,'Totalplay',1,1,'Semestral','Saldo/Plan Telefónico',4444,'2024-06-06'),(4,20,1,1,'Totalplay',1,1,'Cuatrimestral','Saldo/Plan Telefónico',4444,'2024-06-06');
/*!40000 ALTER TABLE `ENCUESTA_SERVICIO` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
/*!50032 DROP TRIGGER IF EXISTS AGREGAR_FECHA_ACTUAL_SERVICIO */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `AGREGAR_FECHA_ACTUAL_SERVICIO` BEFORE INSERT ON `ENCUESTA_SERVICIO` FOR EACH ROW BEGIN
    SET NEW.FECHA = CURDATE();
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
/*!50032 DROP TRIGGER IF EXISTS ACTUALIZAR_FECHA_ACTUAL_SERVICIO */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `ACTUALIZAR_FECHA_ACTUAL_SERVICIO` BEFORE UPDATE ON `ENCUESTA_SERVICIO` FOR EACH ROW BEGIN
    SET NEW.FECHA = CURDATE();
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `ENCUESTA_TRANSPORTE`
--

DROP TABLE IF EXISTS `ENCUESTA_TRANSPORTE`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ENCUESTA_TRANSPORTE` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ID_USUARIO` int DEFAULT NULL,
  `TRANSPORTE_PRINCIPAL` varchar(100) DEFAULT NULL,
  `TRANSPORTE_SECUNDARIO` varchar(100) DEFAULT NULL,
  `FRECUENCIA_USO` varchar(100) DEFAULT NULL,
  `PUNTOS_ACCESIBLES` tinyint(1) DEFAULT NULL,
  `LUGAR_DESTINO_FRECUENTE` varchar(100) DEFAULT NULL,
  `TIEMPO_TRASLADO` varchar(100) DEFAULT NULL,
  `FECHA` date DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_USUARIO` (`ID_USUARIO`),
  CONSTRAINT `ENCUESTA_TRANSPORTE_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ENCUESTA_TRANSPORTE`
--

LOCK TABLES `ENCUESTA_TRANSPORTE` WRITE;
/*!40000 ALTER TABLE `ENCUESTA_TRANSPORTE` DISABLE KEYS */;
INSERT INTO `ENCUESTA_TRANSPORTE` (`ID`, `ID_USUARIO`, `TRANSPORTE_PRINCIPAL`, `TRANSPORTE_SECUNDARIO`, `FRECUENCIA_USO`, `PUNTOS_ACCESIBLES`, `LUGAR_DESTINO_FRECUENTE`, `TIEMPO_TRASLADO`, `FECHA`) VALUES (1,5,'Autobus','Caminar','Quincenal',1,'Escuela','3','2024-06-05'),(2,16,'Motocicleta','Bicicleta','Quincenal',1,'Escuela','5','2024-06-06'),(3,18,'Motocicleta','Bicicleta','Semanal',0,'Escuela','3','2024-06-06'),(4,19,'Autobus','Ruta','Semanal',0,'Escuela','4','2024-06-06'),(5,20,'Ruta','Bicicleta','Semanal',0,'Escuela','3','2024-06-06');
/*!40000 ALTER TABLE `ENCUESTA_TRANSPORTE` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
/*!50032 DROP TRIGGER IF EXISTS AGREGAR_FECHA_ACTUAL_TRANSPORTE */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `AGREGAR_FECHA_ACTUAL_TRANSPORTE` BEFORE INSERT ON `ENCUESTA_TRANSPORTE` FOR EACH ROW BEGIN
    SET NEW.FECHA = CURDATE();
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
/*!50032 DROP TRIGGER IF EXISTS ACTUALIZAR_FECHA_ACTUAL_TRANSPORTE */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `ACTUALIZAR_FECHA_ACTUAL_TRANSPORTE` BEFORE UPDATE ON `ENCUESTA_TRANSPORTE` FOR EACH ROW BEGIN
    SET NEW.FECHA = CURDATE();
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `FORO_PREGUNTA`
--

DROP TABLE IF EXISTS `FORO_PREGUNTA`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `FORO_PREGUNTA` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ID_USUARIO` int DEFAULT NULL,
  `PREGUNTA` varchar(1000) DEFAULT NULL,
  `NOMBRE` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_USUARIO` (`ID_USUARIO`),
  CONSTRAINT `FORO_PREGUNTA_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `FORO_PREGUNTA`
--

LOCK TABLES `FORO_PREGUNTA` WRITE;
/*!40000 ALTER TABLE `FORO_PREGUNTA` DISABLE KEYS */;
/*!40000 ALTER TABLE `FORO_PREGUNTA` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `FORO_PREGUNTA_RESPUESTA`
--

DROP TABLE IF EXISTS `FORO_PREGUNTA_RESPUESTA`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `FORO_PREGUNTA_RESPUESTA` (
  `ID_PREGUNTA` int DEFAULT NULL,
  `ID_RESPUESTA` int DEFAULT NULL,
  KEY `ID_PREGUNTA` (`ID_PREGUNTA`),
  KEY `ID_RESPUESTA` (`ID_RESPUESTA`),
  CONSTRAINT `FORO_PREGUNTA_RESPUESTA_ibfk_1` FOREIGN KEY (`ID_PREGUNTA`) REFERENCES `FORO_PREGUNTA` (`ID`),
  CONSTRAINT `FORO_PREGUNTA_RESPUESTA_ibfk_2` FOREIGN KEY (`ID_RESPUESTA`) REFERENCES `FORO_RESPUESTA` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `FORO_PREGUNTA_RESPUESTA`
--

LOCK TABLES `FORO_PREGUNTA_RESPUESTA` WRITE;
/*!40000 ALTER TABLE `FORO_PREGUNTA_RESPUESTA` DISABLE KEYS */;
/*!40000 ALTER TABLE `FORO_PREGUNTA_RESPUESTA` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `FORO_RESPUESTA`
--

DROP TABLE IF EXISTS `FORO_RESPUESTA`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `FORO_RESPUESTA` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `ID_USUARIO` int DEFAULT NULL,
  `RESPUESTA` varchar(1000) DEFAULT NULL,
  `NOMBRE` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_USUARIO` (`ID_USUARIO`),
  CONSTRAINT `FORO_RESPUESTA_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `FORO_RESPUESTA`
--

LOCK TABLES `FORO_RESPUESTA` WRITE;
/*!40000 ALTER TABLE `FORO_RESPUESTA` DISABLE KEYS */;
/*!40000 ALTER TABLE `FORO_RESPUESTA` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PUBLICACION_EVENTO`
--

DROP TABLE IF EXISTS `PUBLICACION_EVENTO`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `PUBLICACION_EVENTO` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `NOMBRE_EVENTO` varchar(100) DEFAULT NULL,
  `LUGAR` varchar(100) DEFAULT NULL,
  `FECHA` date DEFAULT NULL,
  `HORA` varchar(100) DEFAULT NULL,
  `UBICACION` varchar(100) DEFAULT NULL,
  `DESCRIPCION_EVENTO` varchar(200) DEFAULT NULL,
  `CATEGORIA` varchar(100) DEFAULT NULL,
  `ID_USUARIO` int DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_USUARIO` (`ID_USUARIO`),
  CONSTRAINT `PUBLICACION_EVENTO_ibfk_1` FOREIGN KEY (`ID_USUARIO`) REFERENCES `USUARIO` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PUBLICACION_EVENTO`
--

LOCK TABLES `PUBLICACION_EVENTO` WRITE;
/*!40000 ALTER TABLE `PUBLICACION_EVENTO` DISABLE KEYS */;
/*!40000 ALTER TABLE `PUBLICACION_EVENTO` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `TIPO_USUARIO`
--

DROP TABLE IF EXISTS `TIPO_USUARIO`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `TIPO_USUARIO` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `TIPO_USUARIO` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TIPO_USUARIO`
--

LOCK TABLES `TIPO_USUARIO` WRITE;
/*!40000 ALTER TABLE `TIPO_USUARIO` DISABLE KEYS */;
INSERT INTO `TIPO_USUARIO` (`ID`, `TIPO_USUARIO`) VALUES (1,'ADMIN');
/*!40000 ALTER TABLE `TIPO_USUARIO` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `USUARIO`
--

DROP TABLE IF EXISTS `USUARIO`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `USUARIO` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `NOMBRE` varchar(100) DEFAULT NULL,
  `APELLIDO_PATERNO` varchar(100) DEFAULT NULL,
  `APELLIDO_MATERNO` varchar(100) DEFAULT NULL,
  `CORREO_ELECTRONICO` varchar(100) DEFAULT NULL,
  `NUMERO_TELEFONO` varchar(100) DEFAULT NULL,
  `USUARIO` varchar(100) DEFAULT NULL,
  `ID_TIPO_USUARIO` int DEFAULT NULL,
  `CONTRASENA` varchar(64) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_TIPO_USUARIO` (`ID_TIPO_USUARIO`),
  CONSTRAINT `USUARIO_ibfk_1` FOREIGN KEY (`ID_TIPO_USUARIO`) REFERENCES `TIPO_USUARIO` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `USUARIO`
--

LOCK TABLES `USUARIO` WRITE;
/*!40000 ALTER TABLE `USUARIO` DISABLE KEYS */;
INSERT INTO `USUARIO` (`ID`, `NOMBRE`, `APELLIDO_PATERNO`, `APELLIDO_MATERNO`, `CORREO_ELECTRONICO`, `NUMERO_TELEFONO`, `USUARIO`, `ID_TIPO_USUARIO`, `CONTRASENA`) VALUES (1,'A','B','C','A@KK.COM','7','A1',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(4,'hola','asd','fgre','asd@gmail.com','3252525','asd12',1,'6f5a4fe99d5a6f206fe1f7f4c154b86ec5225b250b1688096a3eedf29591e1d3'),(5,'Tomas','Villa','Alegre','a@gmail.com','7776789843','tommy1',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(7,'1','1','1','1@1.com','1234567890','user1',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(8,'1','1','1','1@1.com','1234567890','user2',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(9,'Christian','Hernandez','Najera','hnco200217@upemor.edu.mx','1234567890','hnco',1,'94f8607915dff25f013e45fc0642fb9830b0fb25ab0ab46d477eaf1061def379'),(10,'Hola2','Hola3','Hola4','hola4@gmail.com','7770985434','Hola4',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(11,'1','1','1','2@2.com','1234567890','12312332131312312',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(13,'Lorena','Perez','H','a@a.com','77791318447','qwe123',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(14,'Axl','Castrejon','Ocampo','r@gmail.com','7771234354','redo',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(16,'Axl','Castrejon','Castrejon','axl@gmail.com','7771234354','Castrejon',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(18,'1','2','3','1@gmail.com','7771234354','Castrejon1',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(19,'1','2','3','2@gmail.com','7771234354','Castrejon1',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3'),(20,'1','2','3','3@gmail.com','7771234354','Castrejon1',1,'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3');
/*!40000 ALTER TABLE `USUARIO` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'ENCUESTA'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-11 13:26:31
