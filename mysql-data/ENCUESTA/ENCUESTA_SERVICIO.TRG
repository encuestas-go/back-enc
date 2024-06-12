TYPE=TRIGGERS
triggers='CREATE DEFINER=`root`@`%` TRIGGER `AGREGAR_FECHA_ACTUAL_SERVICIO` BEFORE INSERT ON `ENCUESTA_SERVICIO` FOR EACH ROW BEGIN\n    SET NEW.FECHA = CURDATE();\nEND' 'CREATE DEFINER=`root`@`%` TRIGGER `ACTUALIZAR_FECHA_ACTUAL_SERVICIO` BEFORE UPDATE ON `ENCUESTA_SERVICIO` FOR EACH ROW BEGIN\n    SET NEW.FECHA = CURDATE();\nEND'
sql_modes=1168113696 1168113696
definers='root@%' 'root@%'
client_cs_names='utf8mb4' 'utf8mb4'
connection_cl_names='utf8mb4_general_ci' 'utf8mb4_general_ci'
db_cl_names='latin1_swedish_ci' 'latin1_swedish_ci'
created=171815968074 171815968082
