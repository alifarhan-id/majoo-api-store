SELECT tgl.tanggal,COALESCE(QRY_TRX.merchant_name,(select  merchants.merchant_name from merchants where merchants.id =1 limit 1 )),
COALESCE(QRY_TRX.omset,0) AS omset 
FROM(	
SELECT generate_series(date '2021-11-1', date '2021-11-30', '1 day') as  Tanggal
) AS tgl 
left join (
select merchants.id, merchants.merchant_name ,date_trunc('day', transactions.created_at) AS tgl_trx, 
sum(transactions.bill_total) as OMSET
from merchants
join transactions on merchants.id = transactions.merchant_id
group by merchants.id, merchants.merchant_name, date_trunc('day', transactions.created_at)
) as QRY_TRX on QRY_TRX.tgl_trx = tgl.Tanggal and QRY_TRX.id =  1

