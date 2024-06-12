TYPE=TRIGGERS
triggers='CREATE DEFINER=`root`@`%` TRIGGER `AGREGAR_FECHA_ACTUAL` BEFORE INSERT ON `ENCUESTA_ACTIVIDAD` FOR EACH ROW BEGIN\n    SET NEW.FECHA = CURDATE();\nEND' 'CREATE DEFINER=`root`@`%` TRIGGER `ACTUALIZAR_FECHA_ACTUAL` BEFORE UPDATE ON `ENCUESTA_ACTIVIDAD` FOR EACH ROW BEGIN\n    SET NEW.FECHA = CURDATE();\nEND'
sql_modes=1168113696 1168113696
definers='root@%' 'root@%'
client_cs_names='utf8mb4' 'utf8mb4'
connection_cl_names='utf8mb4_general_ci' 'utf8mb4_general_ci'
db_cl_names='latin1_swedish_ci' 'latin1_swedish_ci'
created=171815967978 171815967985
