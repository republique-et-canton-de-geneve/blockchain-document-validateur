<div class="<?php print $classes ?>" <?php if (!empty($css_id)) { print "id=\"$css_id\""; } ?>>
  <?php if ($content['top'] || $content['left1'] || $content['right1']): ?>
  <div class="row-introduction">
    <?php if ($content['bottom']): ?>
      <div class="row">
        <?php print $content['top']; ?>
      </div>
    <?php endif; ?>
    <?php if ($content['left1'] || $content['right1']): ?>
      <div class="row">
        <?php print $content['left1']; ?>
        <?php print $content['right1']; ?>
      </div>
    <?php endif; ?>
  </div>
  <?php endif; ?>
  <?php if ($content['left2'] || $content['right2']): ?>
    <div class="row">
      <?php print $content['left2']; ?>
      <?php print $content['right2']; ?>
    </div>
  <?php endif; ?>
  <?php if ($content['left3'] || $content['right3']): ?>
    <div class="row">
      <?php print $content['left3']; ?>
      <?php print $content['right3']; ?>
    </div>
  <?php endif; ?>
  <?php if ($content['bottom']): ?>
    <div class="row">
      <?php print $content['bottom']; ?>
    </div>
  <?php endif; ?>
</div>