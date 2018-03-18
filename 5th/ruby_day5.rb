require 'digest'

s = Time.now
res = []
i = 0
while res.size < 8 do
  dig = Digest::MD5.hexdigest "ffykfhsq#{i}"
  res << dig[5] if dig.start_with?('00000')
  i+=1
end
puts res.join('')
puts Time.now - s
