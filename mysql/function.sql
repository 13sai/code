DROP FUNCTION IF EXISTS sai_filter;
Delimiter //
CREATE 
    FUNCTION sai_filter(ext JSON) 
    RETURNS JSON        
BEGIN
    DECLARE res JSON DEFAULT '[]';
    SET @uni_str = '';
    IF (ext is NULL OR ext = '[]' OR ext = '"[]"')  THEN
        return '[]';
    END IF;

    SET ext = trim(BOTH '"' FROM ext);

    IF JSON_VALID(ext) = 0 THEN
        return '[]';
    END IF;

    SET @json_len = JSON_LENGTH(ext);
    SET @i = 0;
    WHILE @i < @json_len DO
        SET @cur = JSON_EXTRACT(ext, CONCAT('$[',@i,']'));
        IF (JSON_VALID(@cur) = 1) and (JSON_EXTRACT(@cur, '$.company_name') is not NULL) and (JSON_UNQUOTE(JSON_EXTRACT(@cur, '$.company_name')) != '') THEN
            SET @cur_key = CONCAT(LOWER(JSON_UNQUOTE(JSON_EXTRACT(@cur, '$.company_name'))), IFNULL(JSON_UNQUOTE(JSON_EXTRACT(@cur, '$.start_date')), '1900-01-01'), IFNULL(JSON_UNQUOTE(JSON_EXTRACT(@cur, '$.end_date')), '1900-01-01'));
            IF find_in_set(@cur_key, @uni_str) = 0 THEN
                SET @uni_str = CONCAT(@uni_str, @cur_key, ',');
                SET res =  JSON_ARRAY_APPEND(res, '$', JSON_EXTRACT(ext, CONCAT('$[',@i,']')));
            END IF;
        END IF;
        SET @i = @i+1;
    END WHILE;
    RETURN res;
END;
//
Delimiter ;