-- Condition 1: W_YTD= sum(D_YTD)
SELECT /*+ parallel(8) */ * FROM(
    SELECT w.w_id, w.w_ytd, d.sum_d_ytd
    FROM bmsql_warehouse w,
    (SELECT /*+ no_use_px parallel(8) */ d_w_id, sum(d_ytd) sum_d_ytd FROM bmsql_district GROUP BY d_w_id) d 
    WHERE w.w_id= d.d_w_id
) x 
WHERE w_ytd != sum_d_ytd; 


