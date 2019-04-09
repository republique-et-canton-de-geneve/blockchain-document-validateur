<div class="<?php print $classes ?>" <?php if (!empty($css_id)) { print "id=\"$css_id\""; } ?>>
  <?php if ($content['top']): ?>
    <div class="row">
      <?php print $content['top']; ?>
    </div>
  <?php endif; ?>
  <?php if ($content['left1'] || $content['middle1'] || $content['right1']): ?>
    <div class="row">
      <?php print $content['left1']; ?>
      <?php print $content['middle1']; ?>
      <?php print $content['right1']; ?>
    </div>
  <?php endif; ?>
  <?php if ($content['left2'] || $content['middle2'] || $content['right2']): ?>
    <div class="row">
      <?php print $content['left2']; ?>
      <?php print $content['middle2']; ?>
      <?php print $content['right2']; ?>
    </div>
  <?php endif; ?>
  <?php if ($content['left3'] || $content['middle3'] || $content['right3']): ?>
    <div class="row">
      <?php print $content['left3']; ?>
      <?php print $content['middle3']; ?>
      <?php print $content['right4']; ?>
    </div>
  <?php endif; ?>
  <?php if ($content['bottom']): ?>
    <div class="row">
      <?php print $content['bottom']; ?>
    </div>
  <?php endif; ?>
</div>